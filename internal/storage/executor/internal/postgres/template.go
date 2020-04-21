package postgres

import (
	"context"
	"database/sql"

	"go.octolab.org/ecosystem/forma/internal/domain"
	"go.octolab.org/ecosystem/forma/internal/errors"
	"go.octolab.org/ecosystem/forma/internal/storage/query"
	"go.octolab.org/ecosystem/forma/internal/storage/types"
)

// NewTemplateContext TODO issue#173
func NewTemplateContext(ctx context.Context, conn *sql.Conn) templateScope {
	return templateScope{ctx, conn}
}

type templateScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO issue#173
func (scope templateScope) Create(token *types.Token, data query.CreateTemplate) (types.Template, error) {
	entity := types.Template{AccountID: token.User.AccountID, Title: data.Title, Definition: data.Definition}
	encoded := string(entity.Definition)
	q := `INSERT INTO "template" ("id", "account_id", "title", "definition")
	      VALUES (coalesce($1, uuid_generate_v4()), $2, $3, $4)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.AccountID, entity.Title, encoded)
	if err := row.Scan(&entity.ID, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to create a template %q",
			token.UserID, token.User.AccountID, entity.Title)
	}
	return entity, nil
}

// Read TODO issue#173
func (scope templateScope) Read(token *types.Token, data query.ReadTemplate) (types.Template, error) {
	entity, encoded := types.Template{ID: data.ID, AccountID: token.User.AccountID}, ""
	q := `SELECT "title", "definition", "created_at", "updated_at", "deleted_at"
	        FROM "template"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if err := row.Scan(&entity.Title, &encoded, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the template %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	entity.Definition = domain.Template(encoded)
	return entity, nil
}

// ReadByID TODO issue#173
func (scope templateScope) ReadByID(id domain.ID) (types.Template, error) {
	entity, encoded := types.Template{ID: id}, ""
	q := `SELECT "title", "definition", "created_at", "updated_at"
	        FROM "template"
	       WHERE "id" = $1 AND "deleted_at" IS NULL`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID)
	if err := row.Scan(&entity.Title, &encoded, &entity.CreatedAt, &entity.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return entity, errors.NotFound(errors.TemplateNotFoundMessage, err, "the template %q not found", entity.ID)
		}
		return entity, errors.Database(errors.ServerErrorMessage, err, "trying to populate the template %q", entity.ID)
	}
	entity.Definition = domain.Template(encoded)
	return entity, nil
}

// Update TODO issue#173
func (scope templateScope) Update(token *types.Token, data query.UpdateTemplate) (types.Template, error) {
	entity, readErr := scope.Read(token, query.ReadTemplate{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	{
		entity.Title = data.Title
		entity.Definition = data.Definition
	}
	encoded := string(entity.Definition)
	q := `UPDATE "template"
	         SET "title" = $1, "definition" = $2
	       WHERE "id" = $3 AND "account_id" = $4
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.Title, encoded, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.UpdatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to update the template %q",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}

// Delete TODO issue#173
func (scope templateScope) Delete(token *types.Token, data query.DeleteTemplate) (types.Template, error) {
	entity, readErr := scope.Read(token, query.ReadTemplate{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	if data.Permanently {
		q := `DELETE FROM "template" WHERE "id" = $1 AND "account_id" = $2 RETURNING now()`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
		if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
			return entity, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to delete the template %q permanently",
				token.UserID, token.User.AccountID, entity.ID)
		}
		return entity, nil
	}
	q := `UPDATE "template"
	         SET "deleted_at" = now()
	       WHERE "id" = $1 AND "account_id" = $2
	   RETURNING "deleted_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, entity.AccountID)
	if scanErr := row.Scan(&entity.DeletedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"user %q of account %q tried to delete the template %q safely",
			token.UserID, token.User.AccountID, entity.ID)
	}
	return entity, nil
}
