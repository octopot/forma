package server

import (
	"html/template"
	"net/http"
	"time"

	"github.com/kamilsk/form-api/domen"
	"github.com/kamilsk/form-api/server/errors"
	"github.com/kamilsk/form-api/server/middleware"
	"github.com/kamilsk/form-api/static"
	"github.com/kamilsk/form-api/transfer/api/v1"
	"github.com/kamilsk/form-api/transfer/encoding"
)

// New returns new instance of Form API server.
// It can raise the panic if HTML templates are not available.
func New(baseURL, tplPath string, service Service) *Server {
	must := func(base, tpl string) string {
		b, err := static.LoadTemplate(base, tpl)
		if err != nil {
			panic(tpl)
		}
		return string(b)
	}
	return &Server{baseURL: baseURL, service: service, templates: struct {
		errorTpl    *template.Template
		redirectTpl *template.Template
	}{
		errorTpl:    template.Must(template.New("error").Parse(must(tplPath, "error.html"))),
		redirectTpl: template.Must(template.New("error").Parse(must(tplPath, "redirect.html"))),
	}}
}

// Server handles HTTP requests.
type Server struct {
	baseURL   string
	service   Service
	templates struct {
		errorTpl    *template.Template
		redirectTpl *template.Template
	}
}

// GetV1 implements `FormAPI` interface.
func (s *Server) GetV1(rw http.ResponseWriter, req *http.Request) {
	var (
		uuid domen.UUID
		enc  encoding.Generic
	)
	{
		uuid = req.Context().Value(middleware.SchemaKey{}).(domen.UUID)
		enc = req.Context().Value(middleware.EncoderKey{}).(encoding.Generic)
	}
	response := s.service.HandleGetV1(v1.GetRequest{UUID: uuid})
	if response.Error != nil {
		errors.FromGetV1(response.Error).MarshalTo(rw) //nolint: errcheck
		return
	}
	response.Schema.Action = s.baseURL + "api/v1/" + response.Schema.ID
	rw.Header().Set("Content-Type", enc.ContentType())
	rw.WriteHeader(http.StatusOK)
	enc.Encode(response.Schema) //nolint: errcheck
}

// PostV1 implements `FormAPI` interface
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
	if err := req.ParseForm(); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		//nolint: errcheck
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
		//nolint: errcheck
		s.templates.errorTpl.Execute(rw, struct {
			Code     int
			Delay    time.Duration
			Redirect string
		}{http.StatusBadRequest, 5 * time.Second, redirect})
		return
	}
	rw.WriteHeader(http.StatusOK)
	//nolint: errcheck
	s.templates.redirectTpl.Execute(rw, struct {
		Title    string
		Delay    time.Duration
		Redirect string
	}{"Success", 5 * time.Second, redirect})
}
