package grpc

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// ProtectedStorage TODO
type ProtectedStorage interface {
	CreateSchema(ctx context.Context, tokenID domain.ID, data query.CreateSchema) (types.Schema, error)
	ReadSchema(ctx context.Context, tokenID domain.ID, data query.ReadSchema) (types.Schema, error)
	UpdateSchema(ctx context.Context, tokenID domain.ID, data query.UpdateSchema) (types.Schema, error)
	DeleteSchema(ctx context.Context, tokenID domain.ID, data query.DeleteSchema) (types.Schema, error)
	CreateTemplate(ctx context.Context, tokenID domain.ID, data query.CreateTemplate) (types.Template, error)
	ReadTemplate(ctx context.Context, tokenID domain.ID, data query.ReadTemplate) (types.Template, error)
	UpdateTemplate(ctx context.Context, tokenID domain.ID, data query.UpdateTemplate) (types.Template, error)
	DeleteTemplate(ctx context.Context, tokenID domain.ID, data query.DeleteTemplate) (types.Template, error)
	ReadInputByID(ctx context.Context, tokenID, id domain.ID) (types.Input, error)
	ReadInputByFilter(ctx context.Context, tokenID domain.ID, data query.InputFilter) ([]types.Input, error)
}
