package router

import "net/http"

// Server defines the behavior of Form API server.
type Server interface {
	// GetV1 is responsible for `GET /api/v1/{UUID}` request handling.
	GetV1(http.ResponseWriter, *http.Request)
	// PostV1 is responsible for `POST /api/v1/{UUID}` request handling.
	PostV1(http.ResponseWriter, *http.Request)
}
