package v1

import "go.octolab.org/ecosystem/forma/internal/domain"

// GetRequest represents `GET /api/v1/{Schema.ID}` request.
type GetRequest struct {
	ID domain.ID
}

// GetResponse represents `GET /api/v1/{Schema.ID}` response.
type GetResponse struct {
	Error  error
	Schema domain.Schema
}
