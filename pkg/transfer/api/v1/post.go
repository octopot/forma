package v1

import (
	"io"

	"github.com/kamilsk/form-api/pkg/domain"
)

// PostRequest represents `POST /api/v1/{Schema.ID}` request.
type PostRequest struct {
	Context   domain.InputContext
	ID        domain.ID
	InputData domain.InputData
	Output    io.Writer
}

// PostResponse represents `POST /api/v1/{Schema.ID}` response.
type PostResponse struct {
	Error error
	URL   string
}
