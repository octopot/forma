package v1

import (
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// GetRequest represents `GET /api/v1/{UUID}` request.
type GetRequest struct {
	UUID   data.UUID
	Format string
}

// GetResponse represents `GET /api/v1/{UUID}` response.
type GetResponse struct {
	Schema form.Schema
	Error  error
}
