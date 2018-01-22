package v1

import "github.com/kamilsk/form-api/domen"

// GetRequest represents `GET /api/v1/{UUID}` request.
type GetRequest struct {
	UUID domen.UUID
}

// GetResponse represents `GET /api/v1/{UUID}` response.
type GetResponse struct {
	Schema domen.Schema
	Error  error
}
