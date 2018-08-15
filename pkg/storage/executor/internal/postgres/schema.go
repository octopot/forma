package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewSchemaContext TODO
func NewSchemaContext(ctx context.Context, conn *sql.Conn) schemaScope {
	return schemaScope{ctx, conn}
}

type schemaScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO
func (scope schemaScope) Create(token *query.Token, data query.CreateSchema) (query.Schema, error) {
	var entity = query.Schema{
		AccountID:  token.User.AccountID,
		Language:   data.Language,
		Title:      data.Title,
		Definition: data.Definition,
	}
	q := `INSERT INTO "schema" ("account_id", "language", "title", "definition") VALUES ($1, $2, $3, $4)
	      RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.AccountID, entity.Language, entity.Title, entity.Definition)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a schema %q", token.UserID, token.User.AccountID, entity.Title)
	}
	return entity, nil
}

// Read TODO
func (scope schemaScope) Read(token *query.Token, data query.ReadSchema) (query.Schema, error) {
	var entity = query.Schema{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "language", "title", "definition", "created_at", "updated_at", "deleted_at" FROM "schema"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.Language, &entity.Title, &entity.Definition,
		&entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the schema %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// ReadByID TODO
func (scope schemaScope) ReadByID(id string) (query.Schema, error) {
	var entity = query.Schema{ID: id}
	q := `SELECT "language", "title", "definition", "created_at", "updated_at" FROM "schema"
	       WHERE "id" = $1 AND "deleted_at" IS NULL`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID)
	if err := row.Scan(&entity.Language, &entity.Title, &entity.Definition,
		&entity.CreatedAt, &entity.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return entity, errors.NotFound(errors.SchemaNotFoundMessage, err, "the schema %q not found", entity.ID)
		}
		return entity, errors.Database(errors.ServerErrorMessage, err, "trying to populate the schema %q", entity.ID)
	}
	return entity, nil
}

// Update TODO
func (scope schemaScope) Update(token *query.Token, data query.UpdateSchema) (query.Schema, error) {
	entity, err := scope.Read(token, query.ReadSchema{ID: data.ID})
	if err != nil {
		return entity, err
	}
	if data.Language != "" {
		entity.Language = data.Language
	}
	if data.Title != "" {
		entity.Title = data.Title
	}
	if data.Definition != "" {
		entity.Definition = data.Definition
	}
	q := `UPDATE "schema" SET "language" = $1, "title" = $2, "definition" = $3
	       WHERE "id" = $4 AND "account_id" = $5
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.Language, entity.Title, entity.Definition,
		entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the schema %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO
func (scope schemaScope) Delete(token *query.Token, data query.DeleteSchema) (query.Schema, error) {
	if data.Permanently {
		// TODO
	}
	entity, err := scope.Read(token, query.ReadSchema{ID: data.ID})
	if err != nil {
		return entity, err
	}
	q := `UPDATE "schema" SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the schema %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
