package v1

import (
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// PostRequest represents `POST /api/v1/{UUID}` request.
type PostRequest struct {
	UUID data.UUID
	Data map[string][]string
}

// PostResponse represents `POST /api/v1/{UUID}` response.
type PostResponse struct {
	Error  error
	ID     int64
	Schema form.Schema
}
