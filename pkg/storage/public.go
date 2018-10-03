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

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return schema, err
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

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return template, err
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
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	entity, err := storage.exec.InputWriter(ctx, conn).Write(query.WriteInput{SchemaID: schemaID, VerifiedData: verified})
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// LogRequest TODO issue#173
func (storage *Storage) LogRequest(ctx context.Context, input *types.Input, meta domain.InputContext) error {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return err
	}
	defer closer()

	// TODO issue#109
	_, _ = storage.exec.LogWriter(ctx, conn).Write(query.WriteLog{
		SchemaID:   input.SchemaID,
		InputID:    input.ID,
		TemplateID: input.Data.Template(),

		// TODO issue#171
		Identifier:   "10000000-2000-4000-8000-160000000000",
		Code:         201,
		InputContext: meta,
	})
	return nil
}
