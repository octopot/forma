package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewLogContext TODO
func NewLogContext(ctx context.Context, conn *sql.Conn) logScope {
	return logScope{ctx, conn}
}

type logScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Write TODO
func (scope logScope) Write(data query.WriteLog) (query.Log, error) {
	entity := query.Log{
		SchemaID: data.SchemaID, InputID: data.InputID, TemplateID: data.TemplateID,
		Identifier: data.Identifier, Code: data.Code, Context: data.InputContext,
	}
	encoded, encodeErr := json.Marshal(entity.Context)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"trying to marshal context `%#v` of the input %q into JSON",
			entity.Context, entity.InputID)
	}
	q := `INSERT INTO "log" ("account_id", "schema_id", "input_id", "template_id", "identifier", "code", "context")
	           VALUES ((SELECT "account_id" FROM "schema" WHERE "schema_id" = $1), $1, $2, $3, $4, $5, $6)
	        RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q,
		entity.SchemaID, entity.InputID, entity.TemplateID,
		entity.Identifier, entity.Code, encoded)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"trying to insert log `%s` of the input %q",
			encoded, entity.InputID)
	}
	return entity, nil
}
