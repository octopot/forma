package v1

import "github.com/kamilsk/form-api/pkg/domain"

// PostRequest represents `POST /api/v1/{Schema.ID}` request.
type PostRequest struct {
	EncryptedMarker string
	ID              domain.ID
	Data            map[string][]string
}

// PostResponse represents `POST /api/v1/{Schema.ID}` response.
type PostResponse struct {
	EncryptedMarker string
	ID              domain.ID
	Error           error
	Schema          domain.Schema
}
