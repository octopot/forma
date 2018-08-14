package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewSchemaContext TODO
func NewSchemaContext(conn *sql.Conn, ctx context.Context) schema {
	return schema{conn, ctx}
}

type schema struct {
	conn *sql.Conn
	ctx  context.Context
}

// Create TODO
func (s schema) Create(token *query.Token, data query.CreateSchema) (query.Schema, error) {
	var entity = query.Schema{
		AccountID:  token.User.AccountID,
		Language:   data.Language,
		Title:      data.Title,
		Definition: data.Definition,
	}
	q := `INSERT INTO "schema" ("account_id", "language", "title", "definition") VALUES ($1, $2, $3, $4)
	      RETURNING "id", "created_at"`
	row := s.conn.QueryRowContext(s.ctx, q, entity.AccountID, entity.Language, entity.Title, entity.Definition)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a schema %q", token.UserID, token.User.AccountID, entity.Title)
	}
	return entity, nil
}

// Read TODO
func (s schema) Read(token *query.Token, data query.ReadSchema) (query.Schema, error) {
	var entity = query.Schema{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "language", "title", "definition", "created_at", "updated_at", "deleted_at" FROM "schema"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := s.conn.QueryRowContext(s.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.Language, &entity.Title, &entity.Definition,
		&entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the schema %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Update TODO
func (s schema) Update(token *query.Token, data query.UpdateSchema) (query.Schema, error) {
	entity, err := s.Read(token, query.ReadSchema{ID: data.ID})
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
	row := s.conn.QueryRowContext(s.ctx, q, entity.Language, entity.Title, entity.Definition,
		entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the schema %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO
func (s schema) Delete(token *query.Token, data query.DeleteSchema) (query.Schema, error) {
	if data.Permanently {
		// TODO
	}
	entity, err := s.Read(token, query.ReadSchema{ID: data.ID})
	if err != nil {
		return entity, err
	}
	q := `UPDATE "schema" SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := s.conn.QueryRowContext(s.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the schema %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
