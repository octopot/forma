package storage

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

/*
 *
 * Schema
 *
 */

// CreateSchema TODO issue#173
func (storage *Storage) CreateSchema(ctx context.Context, tokenID domain.ID, data query.CreateSchema) (types.Schema, error) {
	var entity types.Schema

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.SchemaEditor(ctx, conn).Create(token, data)
}

// ReadSchema TODO issue#173
func (storage *Storage) ReadSchema(ctx context.Context, tokenID domain.ID, data query.ReadSchema) (types.Schema, error) {
	var entity types.Schema

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.SchemaEditor(ctx, conn).Read(token, data)
}

// UpdateSchema TODO issue#173
func (storage *Storage) UpdateSchema(ctx context.Context, tokenID domain.ID, data query.UpdateSchema) (types.Schema, error) {
	var entity types.Schema

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.SchemaEditor(ctx, conn).Update(token, data)
}

// DeleteSchema TODO issue#173
func (storage *Storage) DeleteSchema(ctx context.Context, tokenID domain.ID, data query.DeleteSchema) (types.Schema, error) {
	var entity types.Schema

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.SchemaEditor(ctx, conn).Delete(token, data)
}

/*
 *
 * Template
 *
 */

// CreateTemplate TODO issue#173
func (storage *Storage) CreateTemplate(ctx context.Context, tokenID domain.ID, data query.CreateTemplate) (types.Template, error) {
	var entity types.Template

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.TemplateEditor(ctx, conn).Create(token, data)
}

// ReadTemplate TODO issue#173
func (storage *Storage) ReadTemplate(ctx context.Context, tokenID domain.ID, data query.ReadTemplate) (types.Template, error) {
	var entity types.Template

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.TemplateEditor(ctx, conn).Read(token, data)
}

// UpdateTemplate TODO issue#173
func (storage *Storage) UpdateTemplate(ctx context.Context, tokenID domain.ID, data query.UpdateTemplate) (types.Template, error) {
	var entity types.Template

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.TemplateEditor(ctx, conn).Update(token, data)
}

// DeleteTemplate TODO issue#173
func (storage *Storage) DeleteTemplate(ctx context.Context, tokenID domain.ID, data query.DeleteTemplate) (types.Template, error) {
	var entity types.Template

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.TemplateEditor(ctx, conn).Delete(token, data)
}

/*
 *
 * Input
 *
 */

// ReadInputByID TODO issue#173
func (storage *Storage) ReadInputByID(ctx context.Context, tokenID domain.ID, id domain.ID) (types.Input, error) {
	var entity types.Input

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return entity, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return entity, err
	}
	return storage.exec.InputReader(ctx, conn).ReadByID(token, id)
}

// ReadInputByFilter TODO issue#173
func (storage *Storage) ReadInputByFilter(ctx context.Context, tokenID domain.ID, filter query.InputFilter) ([]types.Input, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	token, err := storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return nil, err
	}
	return storage.exec.InputReader(ctx, conn).ReadByFilter(token, filter)
}
