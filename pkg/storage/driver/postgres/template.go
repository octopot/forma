package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage"
	"github.com/kamilsk/form-api/pkg/storage/driver"
)

// NewTemplateContext TODO
func NewTemplateContext(conn *sql.Conn, ctx context.Context) (driver.TemplateEditor, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	return template{conn, ctx}, func() {
		cancel()
		_ = conn.Close()
	}
}

type template struct {
	conn *sql.Conn
	ctx  context.Context
}

// Create TODO
func (s template) Create(token *storage.Token, data driver.CreateTemplate) (storage.Template, error) {
	var entity = storage.Template{
		AccountID:  token.User.AccountID,
		Title:      data.Title,
		Definition: data.Definition,
	}
	query := `INSERT INTO "template" ("account_id", "title", "definition") VALUES ($1, $2, $3)
	          RETURNING "id", "created_at"`
	row := s.conn.QueryRowContext(s.ctx, query, entity.AccountID, entity.Title, entity.Definition)
	if row.Scan(&entity.ID, &entity.CreatedAt) != nil {
		return entity, errors.Database(errors.ServerErrorMessage, row.Scan(),
			"user %q of account %q tried to create a template %q", token.UserID, token.User.AccountID, entity.Title)
	}
	return entity, nil
}

// Read TODO
func (s template) Read(token *storage.Token, data driver.ReadTemplate) (storage.Template, error) {
	var entity = storage.Template{ID: data.ID, AccountID: token.User.AccountID}
	query := `SELECT "title", "definition", "created_at", "updated_at", "deleted_at" FROM "template"
	          WHERE "id" = $1 AND "account_id" = $2`
	row := s.conn.QueryRowContext(s.ctx, query, entity.ID, entity.AccountID)
	if row.Scan(&entity.Title, &entity.Definition, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt) != nil {
		return entity, errors.Database(errors.ServerErrorMessage, row.Scan(),
			"user %q of account %q tried to read the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Update TODO
func (s template) Update(token *storage.Token, data driver.UpdateTemplate) (storage.Template, error) {
	entity, err := s.Read(token, driver.ReadTemplate{ID: data.ID})
	if err != nil {
		return entity, err
	}
	if data.Title != "" {
		entity.Title = data.Title
	}
	if data.Definition != "" {
		entity.Definition = data.Definition
	}
	query := `UPDATE "template" SET "title" = $1, "definition" = $2
	          WHERE "id" = $3 AND "account_id" = $4
	          RETURNING "updated_at"`
	row := s.conn.QueryRowContext(s.ctx, query, entity.Title, entity.Definition,
		entity.ID, entity.AccountID)
	if row.Scan(&entity.UpdatedAt) != nil {
		return entity, errors.Database(errors.ServerErrorMessage, row.Scan(),
			"user %q of account %q tried to update the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO
func (s template) Delete(token *storage.Token, data driver.DeleteTemplate) (storage.Template, error) {
	if data.Permanently {
		// TODO
	}
	entity, err := s.Read(token, driver.ReadTemplate{ID: data.ID})
	if err != nil {
		return entity, err
	}
	query := `UPDATE "template" SET "deleted_at" = now()
	          WHERE "id" = $1 AND "account_id" = $2
	          RETURNING "deleted_at"`
	row := s.conn.QueryRowContext(s.ctx, query, entity.ID, entity.AccountID)
	if row.Scan(&entity.DeletedAt) != nil {
		return entity, errors.Database(errors.ServerErrorMessage, row.Scan(),
			"user %q of account %q tried to delete the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
