package v1

import "github.com/kamilsk/form-api/domain"

// PostRequest represents `POST /api/v1/{UUID}` request.
type PostRequest struct {
	UUID domain.UUID
	Data map[string][]string
}

// PostResponse represents `POST /api/v1/{UUID}` response.
type PostResponse struct {
	ID     int64
	Error  error
	Schema domain.Schema
}
