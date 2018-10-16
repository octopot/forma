package server

import (
	"net/http"
	"time"

	deep "github.com/pkg/errors"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/server/middleware"
	"github.com/kamilsk/form-api/pkg/static"
	"github.com/kamilsk/form-api/pkg/transfer/api/v1"
	"github.com/kamilsk/form-api/pkg/transfer/encoding"
)

// New returns a new instance of the Forma server.
func New(cnf config.ServerConfig, service Service) *Server {
	return &Server{cnf, service}
}

// Server handles HTTP requests.
type Server struct {
	config  config.ServerConfig
	service Service
}

// GetV1 is responsible for `GET /api/v1/{Schema.ID}` request handling.
// Deprecated: TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
func (s *Server) GetV1(rw http.ResponseWriter, req *http.Request) {
	var (
		id      = req.Context().Value(middleware.SchemaKey{}).(domain.ID)
		encoder = req.Context().Value(middleware.EncoderKey{}).(encoding.Generic)
	)
	resp := s.service.HandleGetV1(req.Context(), v1.GetRequest{ID: id})
	if resp.Error != nil {
		if err, is := resp.Error.(errors.ApplicationError); is {
			if _, is = err.IsClientError(); is {
				rw.WriteHeader(http.StatusNotFound)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", encoder.ContentType())
	rw.WriteHeader(http.StatusOK)
	encoder.Encode(resp.Schema)
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
		uuid     = req.Context().Value(middleware.SchemaKey{}).(domain.ID)
		request  = v1.PostRequest{ID: uuid, InputData: domain.InputData(req.PostForm), InputContext: domain.InputContext{}}
		response = s.service.HandlePostV1(request)
		redirect = request.InputData.Redirect(req.Referer(), response.Schema.Action)
	)
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
						response.Schema.Method = http.MethodPost
						response.Schema.EncodingType = "application/x-www-form-urlencoded"
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

func must(base, tpl string) string {
	b, err := static.LoadTemplate(base, tpl)
	if err != nil {
		panic(tpl)
	}
	return string(b)
}
