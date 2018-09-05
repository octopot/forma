package grpc

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// ProtectedStorage TODO issue#173
type ProtectedStorage interface {
	// CreateSchema TODO issue#173
	CreateSchema(ctx context.Context, tokenID domain.ID, data query.CreateSchema) (types.Schema, error)
	// ReadSchema TODO issue#173
	ReadSchema(ctx context.Context, tokenID domain.ID, data query.ReadSchema) (types.Schema, error)
	// UpdateSchema TODO issue#173
	UpdateSchema(ctx context.Context, tokenID domain.ID, data query.UpdateSchema) (types.Schema, error)
	// DeleteSchema TODO issue#173
	DeleteSchema(ctx context.Context, tokenID domain.ID, data query.DeleteSchema) (types.Schema, error)
	// CreateTemplate TODO issue#173
	CreateTemplate(ctx context.Context, tokenID domain.ID, data query.CreateTemplate) (types.Template, error)
	// ReadTemplate TODO issue#173
	ReadTemplate(ctx context.Context, tokenID domain.ID, data query.ReadTemplate) (types.Template, error)
	// UpdateTemplate TODO issue#173
	UpdateTemplate(ctx context.Context, tokenID domain.ID, data query.UpdateTemplate) (types.Template, error)
	// DeleteTemplate TODO issue#173
	DeleteTemplate(ctx context.Context, tokenID domain.ID, data query.DeleteTemplate) (types.Template, error)
	// ReadInputByID TODO issue#173
	ReadInputByID(ctx context.Context, tokenID, id domain.ID) (types.Input, error)
	// ReadInputByFilter TODO issue#173
	ReadInputByFilter(ctx context.Context, tokenID domain.ID, data query.InputFilter) ([]types.Input, error)
}
