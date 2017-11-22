package server

import (
	"net/http"

	"github.com/kamilsk/form-api/data/transfer/api/v1"
)

// FormAPI defines behavior of Form API server.
type FormAPI interface {
	// GetV1 responses for `GET /api/v1/{UUID}` request handling.
	GetV1(http.ResponseWriter, *http.Request)
	// PostV1 responses for `POST /api/v1/{UUID}` request handling.
	PostV1(http.ResponseWriter, *http.Request)
}

// FormAPIService defines behavior of Form API service.
type FormAPIService interface {
	// HandleGetV1 handles `GET /api/v1/{UUID}` request.
	HandleGetV1(v1.GetRequest) v1.GetResponse
	// HandlePostV1 handles `POST /api/v1/{UUID}` request.
	HandlePostV1(v1.PostRequest) v1.PostResponse
}
