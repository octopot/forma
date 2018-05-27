package server

import (
	"encoding/base64"
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
	"path"
	"time"

	deep "github.com/pkg/errors"

	"github.com/kamilsk/form-api/domain"
	"github.com/kamilsk/form-api/errors"
	"github.com/kamilsk/form-api/server/middleware"
	"github.com/kamilsk/form-api/static"
	"github.com/kamilsk/form-api/transfer/api/v1"
	"github.com/kamilsk/form-api/transfer/encoding"
)

const (
	tokenCookieName = "token"
	redirectKey     = "_redirect"
	timeoutKey      = "_timeout"
)

// New returns a new instance of Form API server.
// It can raise the panic if base URL is invalid or HTML templates are not available.
func New(baseURL, tplPath string, service Service) *Server {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	return &Server{baseURL: u, service: service, templates: struct {
		errorTpl    *template.Template
		redirectTpl *template.Template
	}{
		errorTpl:    template.Must(template.New("error").Parse(must(tplPath, "error.html"))),
		redirectTpl: template.Must(template.New("redirect").Parse(must(tplPath, "redirect.html"))),
	}}
}

// Server handles HTTP requests.
type Server struct {
	baseURL   *url.URL
	service   Service
	templates struct {
		errorTpl    *template.Template
		redirectTpl *template.Template
	}
}

// GetV1 is responsible for `GET /api/v1/{Schema.ID}` request handling.
func (s *Server) GetV1(rw http.ResponseWriter, req *http.Request) {
	var (
		uuid = req.Context().Value(middleware.SchemaKey{}).(domain.UUID)
		enc  = req.Context().Value(middleware.EncoderKey{}).(encoding.Generic)
	)
	response := s.service.HandleGetV1(v1.GetRequest{UUID: uuid})
	if response.Error != nil {
		if err, is := response.Error.(errors.ApplicationError); is {
			if _, is := err.IsClientError(); is {
				rw.WriteHeader(http.StatusNotFound)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	{ // domain logic
		// add form namespace to make its elements unique
		response.Schema.ID = string(uuid)
		for i := range response.Schema.Inputs {
			response.Schema.Inputs[i].ID = string(uuid) + "_" + response.Schema.Inputs[i].Name
		}
		// replace fallback by current API call
		response.Schema.Action = extend(*s.baseURL, "api/v1", string(uuid))
	}

	rw.Header().Set("Content-Type", enc.ContentType())
	rw.WriteHeader(http.StatusOK)
	enc.Encode(response.Schema)
}

// PostV1 is responsible for `POST /api/v1/{Schema.ID}` request handling.
func (s *Server) PostV1(rw http.ResponseWriter, req *http.Request) {
	type feedback struct {
		ID     string `json:"id"`
		Result string `json:"result"`
	}
	if err := req.ParseForm(); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var (
		uuid = req.Context().Value(middleware.SchemaKey{}).(domain.UUID)
	)

	// TODO: move to middleware layer
	// TODO: support opts.Anonymously()
	cookie, err := req.Cookie(tokenCookieName)
	if err != nil {
		cookie = &http.Cookie{Name: tokenCookieName}
	}

	response := s.service.HandlePostV1(v1.PostRequest{EncryptedMarker: cookie.Value, UUID: uuid, Data: req.PostForm})

	// TODO: move to middleware layer
	// TODO: support opts.Anonymously()
	cookie.MaxAge, cookie.Path, cookie.Value = 0, "/", response.EncryptedMarker
	cookie.Secure, cookie.HttpOnly = true, true
	http.SetCookie(rw, cookie)

	redirect := fallback(req.PostFormValue(redirectKey), req.Referer(), response.Schema.Action)
	if response.Error != nil {
		if err, is := response.Error.(errors.ApplicationError); is {
			if clientErr, is := err.IsClientError(); is {
				switch {
				case clientErr.IsResourceNotFound():
					rw.WriteHeader(http.StatusNotFound)
				case clientErr.IsInvalidInput():

					{ // domain logic
						// add form namespace to make its elements unique
						response.Schema.ID = string(uuid)
						for i := range response.Schema.Inputs {
							response.Schema.Inputs[i].ID = string(uuid) + "_" + response.Schema.Inputs[i].Name
						}
						// replace fallback by current API call
						response.Schema.Action = extend(*s.baseURL, "api/v1", string(uuid))
						// add URL marker
						u, err := url.Parse(redirect)
						if err == nil {
							u.Fragment = base64.StdEncoding.EncodeToString(func() []byte {
								raw, _ := json.Marshal(feedback{ID: string(uuid), Result: "failure"})
								return raw
							}())
							redirect = u.String()
						}
					}

					cause := deep.Cause(err).(domain.ValidationError)
					rw.WriteHeader(http.StatusBadRequest)
					s.templates.errorTpl.Execute(rw, static.ErrorPageContext{
						Schema:   response.Schema,
						Error:    cause,
						Delay:    30 * time.Duration(len(cause.InputWithErrors())) * time.Second,
						Redirect: redirect,
					})
				}
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	{ // domain logic
		// add URL marker
		u, err := url.Parse(redirect)
		if err == nil {
			u.Fragment = base64.StdEncoding.EncodeToString(func() []byte {
				raw, _ := json.Marshal(feedback{ID: string(uuid), Result: "success"})
				return raw
			}())
			redirect = u.String()
		}
	}

	rw.Header().Set("Location", redirect)
	rw.WriteHeader(http.StatusFound)
}

func extend(u url.URL, paths ...string) string {
	if len(paths) == 0 {
		return u.String()
	}
	u.Path = path.Join(append([]string{u.Path}, paths...)...)
	return u.String()
}

func fallback(value string, fallbackValues ...string) string {
	if value == "" {
		for _, value := range fallbackValues {
			if value != "" {
				return value
			}
		}
	}
	return value
}

func must(base, tpl string) string {
	b, err := static.LoadTemplate(base, tpl)
	if err != nil {
		panic(tpl)
	}
	return string(b)
}
