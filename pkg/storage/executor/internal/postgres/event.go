package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// NewEventContext TODO issue#173
func NewEventContext(ctx context.Context, conn *sql.Conn) eventScope {
	return eventScope{ctx, conn}
}

type eventScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Write TODO issue#173
func (scope eventScope) Write(data query.WriteLog) (types.Event, error) {
	entity := types.Event{
		SchemaID:   data.SchemaID,
		InputID:    data.InputID,
		TemplateID: data.TemplateID,
		Identifier: data.Identifier,
		Context:    data.Context,
		Code:       data.Code,
		URL:        data.URL,
	}
	encoded, encodeErr := json.Marshal(entity.Context)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"trying to marshal context `%#v` of the input %q into JSON",
			entity.Context, entity.InputID)
	}
	q := `INSERT INTO "event" ("account_id", "schema_id", "input_id", "template_id", "identifier", "context", "code", "url")
	      VALUES ((SELECT "account_id" FROM "schema" WHERE "id" = $1), $1, $2, $3, $4, $5, $6, $7)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q,
		entity.SchemaID, entity.InputID, entity.TemplateID, entity.Identifier,
		encoded, entity.Code, entity.URL,
	)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"trying to insert event `%s` of the input %q (%+v)",
			encoded, entity.InputID, data)
	}
	return entity, nil
}
