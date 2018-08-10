package postgres

import (
	"bytes"
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage"
	"github.com/kamilsk/form-api/pkg/storage/driver"
)

// NewInputContext TODO
func NewInputContext(conn *sql.Conn, ctx context.Context) (driver.InputReader, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	return input{conn, ctx}, func() {
		cancel()
		_ = conn.Close()
	}
}

type input struct {
	conn *sql.Conn
	ctx  context.Context
}

// DataByID TODO
// TODO check access
func (i input) DataByID(token *storage.Token, id string) (storage.Input, error) {
	var entity = storage.Input{ID: id}
	query := `SELECT "schema_id", "data", "created_at" FROM "input" WHERE "id" = $1`
	row := i.conn.QueryRowContext(i.ctx, query, entity.ID)
	if row.Scan(&entity.SchemaID, &entity.Data, &entity.CreatedAt) != nil {
		return entity, errors.Database(errors.ServerErrorMessage, row.Scan(),
			"user %q of account %q tried to read the input %q", token.UserID, token.User.AccountID, id)
	}
	return entity, nil
}

// DataByFilter TODO
// TODO check access
func (i input) DataByFilter(token *storage.Token, filter driver.InputFilter) ([]storage.Input, error) {
	args := make([]interface{}, 0, 3)
	args = append(args, filter.SchemaID)
	// go 1.10
	// builder := strings.Builder{}
	builder := bytes.NewBuffer(make([]byte, 0, 82+35))
	_, _ = builder.WriteString(`SELECT "id", "schema_id", "data", "created_at" FROM "input" WHERE "schema_id" = $1`)
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
	var entities = make([]storage.Input, 0, 8)
	rows, err := i.conn.QueryContext(i.ctx, builder.String(), args...)
	if err != nil {
		return nil, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read inputs by criteria %+v", token.UserID, token.User.AccountID, filter)
	}
	for rows.Next() {
		var entity storage.Input
		if rows.Scan(&entity.ID, &entity.SchemaID, &entity.Data, &entity.CreatedAt) != nil {
			return nil, errors.Database(errors.ServerErrorMessage, err,
				"user %q of account %q tried to read inputs by criteria %+v", token.UserID, token.User.AccountID, filter)
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
