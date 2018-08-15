package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewInputContext TODO
func NewInputContext(conn *sql.Conn, ctx context.Context) inputScope {
	return inputScope{conn, ctx}
}

type inputScope struct {
	conn *sql.Conn
	ctx  context.Context
}

// Write TODO
func (scope inputScope) Write(data query.WriteInput) (query.Input, error) {
	var entity = query.Input{SchemaID: data.SchemaID}
	encoded, err := json.Marshal(data.VerifiedData)
	if err != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, err,
			"trying to marshal data `%#v` for the schema %q into JSON", data.VerifiedData, data.SchemaID)
	}
	entity.Data = encoded
	q := `INSERT INTO "input" ("schema_id", "data") VALUES ($1, $2) RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.SchemaID, entity.Data)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"trying to insert input `%s` for the schema %q", entity.Data, entity.SchemaID)
	}
	return entity, nil
}

// ReadByID TODO
// TODO check access
func (scope inputScope) ReadByID(token *query.Token, id string) (query.Input, error) {
	var entity = query.Input{ID: id}
	q := `SELECT "schema_id", "data", "created_at" FROM "input" WHERE "id" = $1`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID)
	if err := row.Scan(&entity.SchemaID, &entity.Data, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the inputScope %q", token.UserID, token.User.AccountID, id)
	}
	return entity, nil
}

// ReadByFilter TODO
// TODO check access
func (scope inputScope) ReadByFilter(token *query.Token, filter query.InputFilter) ([]query.Input, error) {
	args := make([]interface{}, 0, 3)
	args = append(args, filter.SchemaID)
	// go 1.10
	// builder := strings.Builder{}
	builder := bytes.NewBuffer(make([]byte, 0, 82+35))
	_, _ = builder.WriteString(`SELECT "id", "data", "created_at" FROM "input" WHERE "schema_id" = $1`)
	switch {
	case !filter.From.IsZero() && !filter.To.IsZero():
		_, _ = builder.WriteString(` AND "created_at" BETWEEN $2 AND $3`)
		args = append(args, filter.From, filter.To)
	case !filter.From.IsZero():
		_, _ = builder.WriteString(` AND "created_at" >= $2`)
		args = append(args, filter.From)
	case !filter.To.IsZero():
		_, _ = builder.WriteString(` AND "created_at" <= $2`)
		args = append(args, filter.To)
	}
	var entities = make([]query.Input, 0, 8)
	rows, err := scope.conn.QueryContext(scope.ctx, builder.String(), args...)
	if err != nil {
		return nil, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read inputs by criteria %+v", token.UserID, token.User.AccountID, filter)
	}
	for rows.Next() {
		var entity = query.Input{SchemaID: filter.SchemaID}
		if scanErr := rows.Scan(&entity.ID, &entity.Data, &entity.CreatedAt); scanErr != nil {
			return nil, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to read inputs by criteria %+v", token.UserID, token.User.AccountID, filter)
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
