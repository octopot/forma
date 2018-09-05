package storage

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// TokenByID TODO
func (storage *Storage) TokenByID(ctx context.Context, id domain.ID) (*types.Token, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	return storage.exec.UserManager(ctx, conn).Token(id)
}

// CreateSchema TODO
func (storage *Storage) CreateSchema(ctx context.Context, token *types.Token, data query.CreateSchema) (types.Schema, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return types.Schema{}, err
	}
	defer closer()

	return storage.exec.SchemaEditor(ctx, conn).Create(token, data)
}

// ReadSchema TODO
func (storage *Storage) ReadSchema(ctx context.Context, token *types.Token, data query.ReadSchema) (types.Schema, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return types.Schema{}, err
	}
	defer closer()

	return storage.exec.SchemaEditor(ctx, conn).Read(token, data)
}

// UpdateSchema TODO
func (storage *Storage) UpdateSchema(ctx context.Context, token *types.Token, data query.UpdateSchema) (types.Schema, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return types.Schema{}, err
	}
	defer closer()

	return storage.exec.SchemaEditor(ctx, conn).Update(token, data)
}

// DeleteSchema TODO
func (storage *Storage) DeleteSchema(ctx context.Context, token *types.Token, data query.DeleteSchema) (types.Schema, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return types.Schema{}, err
	}
	defer closer()

	return storage.exec.SchemaEditor(ctx, conn).Delete(token, data)
}

// CreateTemplate TODO
func (storage *Storage) CreateTemplate(ctx context.Context, token *types.Token, data query.CreateTemplate) (types.Template, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return types.Template{}, err
	}
	defer closer()

	return storage.exec.TemplateEditor(ctx, conn).Create(token, data)
}

// ReadTemplate TODO
func (storage *Storage) ReadTemplate(ctx context.Context, token *types.Token, data query.ReadTemplate) (types.Template, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return types.Template{}, err
	}
	defer closer()

	return storage.exec.TemplateEditor(ctx, conn).Read(token, data)
}

// UpdateTemplate TODO
func (storage *Storage) UpdateTemplate(ctx context.Context, token *types.Token, data query.UpdateTemplate) (types.Template, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return types.Template{}, err
	}
	defer closer()

	return storage.exec.TemplateEditor(ctx, conn).Update(token, data)
}

// DeleteTemplate TODO
func (storage *Storage) DeleteTemplate(ctx context.Context, token *types.Token, data query.DeleteTemplate) (types.Template, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return types.Template{}, err
	}
	defer closer()

	return storage.exec.TemplateEditor(ctx, conn).Delete(token, data)
}
