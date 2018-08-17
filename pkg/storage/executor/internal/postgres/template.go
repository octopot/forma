package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewTemplateContext TODO
func NewTemplateContext(ctx context.Context, conn *sql.Conn) templateScope {
	return templateScope{ctx, conn}
}

type templateScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO
func (scope templateScope) Create(token *query.Token, data query.CreateTemplate) (query.Template, error) {
	var entity = query.Template{
		AccountID:  token.User.AccountID,
		Title:      data.Title,
		Definition: data.Definition,
	}
	q := `INSERT INTO "template" ("account_id", "title", "definition") VALUES ($1, $2, $3)
	      RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.AccountID, entity.Title, entity.Definition)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a template %q", token.UserID, token.User.AccountID, entity.Title)
	}
	return entity, nil
}

// Read TODO
func (scope templateScope) Read(token *query.Token, data query.ReadTemplate) (query.Template, error) {
	var entity = query.Template{ID: data.ID, AccountID: token.User.AccountID}
	q := `SELECT "title", "definition", "created_at", "updated_at", "deleted_at" FROM "template"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.Title, &entity.Definition,
		&entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// ReadByID TODO
func (scope templateScope) ReadByID(id domain.ID) (query.Template, error) {
	var entity = query.Template{ID: id}
	q := `SELECT "title", "definition", "created_at", "updated_at" FROM "template"
	       WHERE "id" = $1 AND "deleted_at" IS NULL`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID)
	if err := row.Scan(&entity.Title, &entity.Definition, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return entity, errors.NotFound(errors.TemplateNotFoundMessage, err, "the template %q not found", entity.ID)
		}
		return entity, errors.Database(errors.ServerErrorMessage, err, "trying to populate the template %q", entity.ID)
	}
	return entity, nil
}

// Update TODO
func (scope templateScope) Update(token *query.Token, data query.UpdateTemplate) (query.Template, error) {
	entity, err := scope.Read(token, query.ReadTemplate{ID: data.ID})
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
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.Title, entity.Definition, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO
func (scope templateScope) Delete(token *query.Token, data query.DeleteTemplate) (query.Template, error) {
	if data.Permanently {
		// TODO
	}
	entity, err := scope.Read(token, query.ReadTemplate{ID: data.ID})
	if err != nil {
		return entity, err
	}
	q := `UPDATE "template" SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the template %q", token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
