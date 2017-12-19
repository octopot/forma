package server

import (
	"html/template"
	"net/http"
	"time"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/encoder"
	"github.com/kamilsk/form-api/data/transfer/api/v1"
	"github.com/kamilsk/form-api/server/errors"
	"github.com/kamilsk/form-api/static"
)

// checklist:
// - inject redirect link via middleware
// - check that Referer exists and it is valid URL
// - classify error: ApplicationError and ValidationError
// - use detailed information in case with ValidationError
// - review templates:
//   - add button next
//   - add bootstrap css
//   - log errors
// - nolint replace by logging

type (
	// EncoderKey used as a context key to store a form schema encoder.
	EncoderKey struct{}
	// UUIDKey used as a context key to store a form schema UUID.
	UUIDKey struct{}
)

// New returns new instance of Form API server.
// It can raise the panic if HTML templates are not available.
func New(baseURL, tplPath string, service FormAPIService) *Server {
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
	service   FormAPIService
	templates struct {
		errorTpl    *template.Template
		redirectTpl *template.Template
	}
}

// GetV1 implements `FormAPI` interface.
func (s *Server) GetV1(rw http.ResponseWriter, req *http.Request) {
	var (
		uuid data.UUID
		enc  encoder.Generic
	)
	{ // from middleware
		uuid = req.Context().Value(UUIDKey{}).(data.UUID)
		enc = req.Context().Value(EncoderKey{}).(encoder.Generic)
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
		uuid     data.UUID
		redirect string
	)
	{ // from middleware
		uuid = req.Context().Value(UUIDKey{}).(data.UUID)
		redirect = req.Header.Get("Referer")
	}

	rw.Header().Set("Content-Type", encoder.HTML)
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
