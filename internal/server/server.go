package server

import (
	"bytes"
	"io"
	"net/http"

	"go.octolab.org/ecosystem/forma/internal/config"
	"go.octolab.org/ecosystem/forma/internal/domain"
	"go.octolab.org/ecosystem/forma/internal/errors"
	"go.octolab.org/ecosystem/forma/internal/server/middleware"
	v1 "go.octolab.org/ecosystem/forma/internal/transfer/api/v1"
	"go.octolab.org/ecosystem/forma/internal/transfer/encoding"
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
	if err := req.ParseForm(); err != nil {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var (
		id      = domain.ID(req.Form.Get("id"))
		encoder = req.Context().Value(middleware.EncoderKey{}).(encoding.Generic)
	)
	if !id.IsValid() {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
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
	_ = encoder.Encode(resp.Schema)
}

// Input is responsible for `POST /api/v1/{Schema.ID}` request handling.
func (s *Server) Input(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	output := bytes.NewBuffer(make([]byte, 0, 1024))
	resp := s.service.HandleInput(req.Context(), v1.PostRequest{
		Context: domain.InputContext{
			Cookies: domain.FromCookies(req.Cookies()),
			Headers: domain.FromHeaders(req.Header),
			Queries: domain.FromRequest(req),
		},
		ID:        domain.ID(req.Form.Get("id")),
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
				_, _ = io.Copy(rw, output)
				return
			}
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Location", resp.URL)
	rw.WriteHeader(http.StatusFound)
}
