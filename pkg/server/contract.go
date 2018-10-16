package server

import (
	"context"

	"github.com/kamilsk/form-api/pkg/transfer/api/v1"
)

// Service defines the behavior of the Forma service.
type Service interface {
	// HandleGetV1 handles an input request.
	// Deprecated: TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
	HandleGetV1(context.Context, v1.GetRequest) v1.GetResponse
	// HandlePostV1 handles an input request.
	HandlePostV1(context.Context, v1.PostRequest) v1.PostResponse
}
