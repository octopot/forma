package storage

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// Schema returns the form schema by provided ID.
func (storage *Storage) Schema(ctx context.Context, id domain.ID) (domain.Schema, error) {
	var schema domain.Schema

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return schema, connErr
	}
	defer closer()

	entity, err := storage.exec.SchemaReader(ctx, conn).ReadByID(id)
	if err != nil {
		return schema, err
	}
	entity.Definition.Title = entity.Title

	return entity.Definition, nil
}

// Template returns the form template by provided ID.
func (storage *Storage) Template(ctx context.Context, id domain.ID) (domain.Template, error) {
	var template domain.Template

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return template, connErr
	}
	defer closer()

	entity, err := storage.exec.TemplateReader(ctx, conn).ReadByID(id)
	if err != nil {
		return template, err
	}
	return entity.Definition, nil
}

// HandleInput TODO issue#173
func (storage *Storage) HandleInput(ctx context.Context, schemaID domain.ID, verified domain.InputData) (*types.Input, error) {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return nil, connErr
	}
	defer closer()

	entity, err := storage.exec.InputWriter(ctx, conn).Write(query.WriteInput{SchemaID: schemaID, VerifiedData: verified})
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// LogInput stores an input event.
func (storage *Storage) LogInput(ctx context.Context, event domain.InputEvent) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer closer()

	// TODO issue#51
	_, writeErr := storage.exec.LogWriter(ctx, conn).Write(query.WriteLog{InputEvent: event})

	return writeErr
}
