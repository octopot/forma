package grpc

import (
	"context"

	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// ProtectedStorage TODO
type ProtectedStorage interface {
	CreateSchema(context.Context, *types.Token, query.CreateSchema) (types.Schema, error)
	ReadSchema(context.Context, *types.Token, query.ReadSchema) (types.Schema, error)
	UpdateSchema(context.Context, *types.Token, query.UpdateSchema) (types.Schema, error)
	DeleteSchema(context.Context, *types.Token, query.DeleteSchema) (types.Schema, error)
}
