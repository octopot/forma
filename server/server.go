package server

import (
	"html/template"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/kamilsk/form-api/domain"
	"github.com/kamilsk/form-api/errors"
	"github.com/kamilsk/form-api/server/middleware"
	"github.com/kamilsk/form-api/static"
	"github.com/kamilsk/form-api/transfer/api/v1"
	"github.com/kamilsk/form-api/transfer/encoding"
)

// New returns new instance of Form API server.
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

// GetV1 is responsible for `GET /api/v1/{UUID}` request handling.
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
		response.Schema.Action = join(*s.baseURL, "api/v1", string(uuid))
	}
	rw.Header().Set("Content-Type", enc.ContentType())
	rw.WriteHeader(http.StatusOK)
	enc.Encode(response.Schema)
}

// PostV1 is responsible for `POST /api/v1/{UUID}` request handling.
func (s *Server) PostV1(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var (
		uuid     = req.Context().Value(middleware.SchemaKey{}).(domain.UUID)
		redirect = req.Header.Get("Referer") // referer, _redirect, action => action fix in data
	)

	rw.Header().Set("Content-Type", encoding.HTML)

	// application/x-www-form-urlencoded
	// application/x-www-form-urlencoded; charset=UTF-8

	request := v1.PostRequest{UUID: uuid, Data: req.PostForm}
	response := s.service.HandlePostV1(request)
	if response.Error != nil {
		rw.WriteHeader(http.StatusBadRequest)
		s.templates.errorTpl.Execute(rw, struct {
			Code     int
			Delay    time.Duration
			Redirect string
		}{http.StatusBadRequest, 5 * time.Second, redirect})
		return
	}
	rw.WriteHeader(http.StatusOK)
	s.templates.redirectTpl.Execute(rw, struct {
		Title    string
		Delay    time.Duration
		Redirect string
	}{"Success", 5 * time.Second, redirect})
}

func join(u url.URL, paths ...string) string {
	if len(paths) == 0 {
		return u.String()
	}
	u.Path = path.Join(append([]string{u.Path}, paths...)...)
	return u.String()
}

func must(base, tpl string) string {
	b, err := static.LoadTemplate(base, tpl)
	if err != nil {
		panic(tpl)
	}
	return string(b)
}
