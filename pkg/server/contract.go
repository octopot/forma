package server

import "github.com/kamilsk/form-api/pkg/transfer/api/v1"

// Service defines the behavior of the Forma service.
type Service interface {
	// HandleGetV1 handles an input request.
	HandleGetV1(v1.GetRequest) v1.GetResponse
	// HandlePostV1 handles an input request.
	HandlePostV1(v1.PostRequest) v1.PostResponse
}
