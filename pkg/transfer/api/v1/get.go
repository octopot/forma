package v1

import "github.com/kamilsk/form-api/pkg/domain"

// GetRequest represents `GET /api/v1/{Schema.ID}` request.
type GetRequest struct {
	ID domain.ID
}

// GetResponse represents `GET /api/v1/{Schema.ID}` response.
type GetResponse struct {
	Schema domain.Schema
	Error  error
}
