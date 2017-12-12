package server

import (
	"net/http"
	"net/url"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/encoder"
	"github.com/kamilsk/form-api/data/transfer/api/v1"
	"github.com/kamilsk/form-api/server/errors"
)

// UUIDKey used as a context key to store a form schema UUID.
type UUIDKey struct{}

// EncoderKey used as sa context key to store a form schema encoder.
type EncoderKey struct{}

// New returns new instance of Form API server.
func New(service FormAPIService) *Server {
	return &Server{service: service}
}

// Server handles HTTP requests.
type Server struct {
	service FormAPIService
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
	rw.Header().Set("Content-Type", enc.ContentType())
	rw.WriteHeader(http.StatusOK)
	enc.Encode(response.Schema) //nolint: errcheck
}

// PostV1 implements `FormAPI` interface
func (s *Server) PostV1(rw http.ResponseWriter, req *http.Request) {
	var uuid data.UUID
	{ // from middleware
		uuid = req.Context().Value(UUIDKey{}).(data.UUID)
	}

	if err := req.ParseForm(); err != nil {
		//httpErr.InvalidFormData(err).MarshalTo(rw) //nolint: errcheck
		return
	}

	referer := req.Header.Get("Referer")
	if referer == "" {
		//httpErr.NoReferer().MarshalTo(rw) //nolint: errcheck
		return
	}

	redirect, err := url.Parse(referer)
	if err != nil {
		//httpErr.InvalidReferer(err).MarshalTo(rw) //nolint: errcheck
		return
	}

	request := v1.PostRequest{UUID: uuid, Data: req.PostForm}
	_ = s.service.HandlePostV1(request)
	rw.Header().Set("Location", redirect.String())
	rw.WriteHeader(http.StatusFound)
}
