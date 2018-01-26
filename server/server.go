package server

import (
	"html/template"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/kamilsk/form-api/domen"
	"github.com/kamilsk/form-api/errors"
	"github.com/kamilsk/form-api/server/middleware"
	"github.com/kamilsk/form-api/static"
	"github.com/kamilsk/form-api/transfer/api/v1"
	"github.com/kamilsk/form-api/transfer/encoding"
)

// New returns new instance of Form API server.
// It can raise the panic if baseURL is invalid or HTML templates are not available.
func New(baseURL, tplPath string, service Service) *Server {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	must := func(base, tpl string) string {
		b, err := static.LoadTemplate(base, tpl)
		if err != nil {
			panic(tpl)
		}
		return string(b)
	}
	return &Server{baseURL: u, service: service, templates: struct {
		errorTpl    *template.Template
		redirectTpl *template.Template
	}{
		errorTpl:    template.Must(template.New("error").Parse(must(tplPath, "error.html"))),
		redirectTpl: template.Must(template.New("error").Parse(must(tplPath, "redirect.html"))),
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
		uuid = req.Context().Value(middleware.SchemaKey{}).(domen.UUID)
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
	response.Schema.Action = join(*s.baseURL, "api/v1", response.Schema.ID)
	rw.Header().Set("Content-Type", enc.ContentType())
	rw.WriteHeader(http.StatusOK)
	enc.Encode(response.Schema)
}

// PostV1 is responsible for `POST /api/v1/{UUID}` request handling.
func (s *Server) PostV1(rw http.ResponseWriter, req *http.Request) {
	var (
		uuid     domen.UUID
		redirect string
	)
	{
		uuid = req.Context().Value(middleware.SchemaKey{}).(domen.UUID)
		redirect = req.Header.Get("Referer")
	}

	rw.Header().Set("Content-Type", encoding.HTML)

	// application/x-www-form-urlencoded
	// application/x-www-form-urlencoded; charset=UTF-8

	if err := req.ParseForm(); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		s.templates.errorTpl.Execute(rw, struct {
			Code     int
			Delay    time.Duration
			Redirect string
		}{http.StatusBadRequest, 5 * time.Second, redirect})
		return
	}

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
