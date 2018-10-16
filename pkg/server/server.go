package server

import (
	"bytes"
	"io"
	"net/http"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/server/middleware"
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

// HandleInput is responsible for `POST /api/v1/{Schema.ID}` request handling.
func (s *Server) HandleInput(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	output := bytes.NewBuffer(make([]byte, 0, 1024))
	resp := s.service.HandlePostV1(req.Context(), v1.PostRequest{
		Context: domain.InputContext{
			Cookies: domain.FromCookies(req.Cookies()),
			Headers: domain.FromHeaders(req.Header),
			Queries: domain.FromRequest(req),
		},
		ID:        req.Context().Value(middleware.SchemaKey{}).(domain.ID),
		InputData: domain.InputData(req.PostForm),
		Output:    output,
	})
	if resp.Error != nil {
		if err, is := resp.Error.(errors.ApplicationError); is {
			if clientErr, isClient := err.IsClientError(); isClient {
				switch {
				case clientErr.IsResourceNotFound():
					rw.WriteHeader(http.StatusNotFound)
				default:
					rw.WriteHeader(http.StatusBadRequest)
				}
				io.Copy(rw, output)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Location", resp.URL)
	rw.WriteHeader(http.StatusFound)
}
