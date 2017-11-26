package server

import (
	"net/http"
	"net/url"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/transfer/api/v1"
)

// UUIDKey used as a context key to store a form UUID.
type UUIDKey struct{}

// New returns new instance of Form API server.
func New(service FormAPIService) *server {
	return &server{service: service}
}

type server struct {
	service FormAPIService
}

// GetV1 implements `FormAPI` interface.
func (s *server) GetV1(rw http.ResponseWriter, req *http.Request) {
	var httpErr *Error

	uuid := req.Context().Value(UUIDKey{}).(data.UUID)

	request := v1.GetRequest{UUID: uuid, Format: req.Header.Get("Accept")}
	response := s.service.HandleGetV1(request)
	if response.Error != nil {
		httpErr.FromGetV1(response.Error).MarshalTo(rw)
		return
	}
	response.Schema.MarshalTo(rw)
}

// PostV1 implements `FormAPI` interface
func (s *server) PostV1(rw http.ResponseWriter, req *http.Request) {
	var httpErr *Error

	uuid := req.Context().Value(UUIDKey{}).(data.UUID)

	if err := req.ParseForm(); err != nil {
		httpErr.InvalidFormData(err).MarshalTo(rw)
		return
	}

	referer := req.Header.Get("Referer")
	if referer == "" {
		httpErr.NoReferer().MarshalTo(rw)
		return
	}

	redirect, err := url.Parse(referer)
	if err != nil {
		httpErr.InvalidReferer(err).MarshalTo(rw)
		return
	}

	request := v1.PostRequest{UUID: uuid, Data: req.PostForm}
	_ = s.service.HandlePostV1(request)
	rw.Header().Set("Location", redirect.String())
	rw.WriteHeader(http.StatusFound)
}
