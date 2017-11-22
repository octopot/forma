package v1

import (
	"net/url"

	"github.com/kamilsk/form-api/data"
)

// PostRequest represents `POST /api/v1/{UUID}` request.
type PostRequest struct {
	UUID data.UUID
	Data url.Values
}

// PostResponse represents `POST /api/v1/{UUID}` response.
type PostResponse struct {
	DefaultRedirect *url.URL
}
