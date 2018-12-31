package server

import (
	"context"

	v1 "github.com/kamilsk/form-api/pkg/transfer/api/v1"
)

// Service defines the behavior of the Forma service.
type Service interface {
	// HandleGetV1 handles an input request.
	// Deprecated: TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
	HandleGetV1(context.Context, v1.GetRequest) v1.GetResponse
	// HandleInput handles an input request.
	HandleInput(context.Context, v1.PostRequest) v1.PostResponse
}
