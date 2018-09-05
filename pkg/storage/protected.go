package storage

import (
	"context"

	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

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
