package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/executor"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewTemplateContext TODO
func NewTemplateContext(conn *sql.Conn, ctx context.Context) executor.TemplateEditor {
	return template{conn, ctx}
}

type template struct {
	conn *sql.Conn
	ctx  context.Context
}

// Create TODO
func (t template) Create(token *query.Token, data query.CreateTemplate) (query.Template, error) {
	var entity = query.Template{
		AccountID:  token.User.AccountID,
		Title:      data.Title,
		Definition: data.Definition,
	}
	q := `INSERT INTO "template" ("account_id", "title", "definition") VALUES ($1, $2, $3)
	      RETURNING "id", "created_at"`
	row := t.conn.QueryRowContext(t.ctx, q, entity.AccountID, entity.Title, entity.Definition)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a template %q", token.UserID, token.User.AccountID, entity.Title)
	}
	return entity, nil
}

// Read TODO
func (t template) Read(token *query.Token, data query.ReadTemplate) (query.Template, error) {
	var entity = query.Template{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "title", "definition", "created_at", "updated_at", "deleted_at" FROM "template"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := t.conn.QueryRowContext(t.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.Title, &entity.Definition,
		&entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Update TODO
func (t template) Update(token *query.Token, data query.UpdateTemplate) (query.Template, error) {
	entity, err := t.Read(token, query.ReadTemplate{ID: data.ID})
	if err != nil {
		return entity, err
	}
	if data.Title != "" {
		entity.Title = data.Title
	}
	if data.Definition != "" {
		entity.Definition = data.Definition
	}
	q := `UPDATE "template" SET "title" = $1, "definition" = $2
	       WHERE "id" = $3 AND "account_id" = $4
	   RETURNING "updated_at"`
	row := t.conn.QueryRowContext(t.ctx, q, entity.Title, entity.Definition,
		entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO
func (t template) Delete(token *query.Token, data query.DeleteTemplate) (query.Template, error) {
	if data.Permanently {
		// TODO
	}
	entity, err := t.Read(token, query.ReadTemplate{ID: data.ID})
	if err != nil {
		return entity, err
	}
	q := `UPDATE "template" SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := t.conn.QueryRowContext(t.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
