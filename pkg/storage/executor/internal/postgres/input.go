package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// NewInputContext TODO issue#173
func NewInputContext(ctx context.Context, conn *sql.Conn) inputScope {
	return inputScope{ctx, conn}
}

type inputScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Write TODO issue#173
func (scope inputScope) Write(data query.WriteInput) (types.Input, error) {
	entity := types.Input{SchemaID: data.SchemaID, Data: data.VerifiedData}
	encoded, encodeErr := json.Marshal(entity.Data)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"trying to marshal data `%#v` for the schema %q into JSON",
			entity.Data, entity.SchemaID)
	}
	q := `INSERT INTO "input" ("schema_id", "data") VALUES ($1, $2) RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.SchemaID, encoded)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"trying to insert input `%s` for the schema %q",
			encoded, entity.SchemaID)
	}
	return entity, nil
}

// ReadByID TODO issue#173
func (scope inputScope) ReadByID(token *types.Token, id domain.ID) (types.Input, error) {
	entity, encoded := types.Input{ID: id}, []byte(nil)
	q := `SELECT "i"."schema_id", "i"."data", "i"."created_at"
	        FROM "input" "i"
	  INNER JOIN "schema" "s" ON "s"."id" = "i"."schema_id"
	       WHERE "i"."id" = $1 AND "s"."account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, token.User.AccountID)
	if err := row.Scan(&entity.SchemaID, &encoded, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the input %q",
			token.UserID, token.User.AccountID, id)
	}
	if err := json.Unmarshal(encoded, &entity.Data); err != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, err,
			"user %q of account %q tried to unmarshal JSON `%s` of the input %q",
			token.UserID, token.User.AccountID, encoded, id)
	}
	return entity, nil
}

// ReadByFilter TODO issue#173
func (scope inputScope) ReadByFilter(token *types.Token, filter query.InputFilter) ([]types.Input, error) {
	q := `SELECT "i"."id", "i"."data", "i"."created_at"
	        FROM "input" "i"
	  INNER JOIN "schema" "s" ON "s"."id" = "i"."schema_id"
	       WHERE "i"."schema_id" = $1 AND "s"."account_id"`
	args := append(make([]interface{}, 0, 4), filter.SchemaID, token.User.AccountID)
	// TODO go 1.10 builder := strings.Builder{}
	builder := bytes.NewBuffer(make([]byte, 0, len(q)+39))
	_, _ = builder.WriteString(q)
	switch {
	case !filter.From.IsZero() && !filter.To.IsZero():
		_, _ = builder.WriteString(` AND "i"."created_at" BETWEEN $2 AND $3`)
		args = append(args, filter.From, filter.To)
	case !filter.From.IsZero():
		_, _ = builder.WriteString(` AND "i"."created_at" >= $2`)
		args = append(args, filter.From)
	case !filter.To.IsZero():
		_, _ = builder.WriteString(` AND "i"."created_at" <= $2`)
		args = append(args, filter.To)
	}
	entities := make([]types.Input, 0, 8)
	rows, dbErr := scope.conn.QueryContext(scope.ctx, builder.String(), args...)
	if dbErr != nil {
		return nil, errors.Database(errors.ServerErrorMessage, dbErr,
			"user %q of account %q tried to read inputs by criteria %+v",
			token.UserID, token.User.AccountID, filter)
	}
	for rows.Next() {
		entity, encoded := types.Input{SchemaID: filter.SchemaID}, []byte(nil)
		if scanErr := rows.Scan(&entity.ID, &encoded, &entity.CreatedAt); scanErr != nil {
			return nil, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to read inputs by criteria %+v",
				token.UserID, token.User.AccountID, filter)
		}
		if decodeErr := json.Unmarshal(encoded, &entity.Data); decodeErr != nil {
			return nil, errors.Serialization(errors.ServerErrorMessage, decodeErr,
				"user %q of account %q tried to unmarshal JSON `%s` of the input %q",
				token.UserID, token.User.AccountID, encoded, entity.ID)
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
