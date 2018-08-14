package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewUserContext TODO
func NewUserContext(conn *sql.Conn, ctx context.Context) manager {
	return manager{conn, ctx}
}

type manager struct {
	conn *sql.Conn
	ctx  context.Context
}

// Token TODO
func (m manager) Token(id string) (*query.Token, error) {
	var (
		token   = query.Token{ID: id}
		user    = query.User{}
		account = query.Account{}
	)
	q := `SELECT "t"."user_id", "t"."expired_at", "t"."created_at",
	             "u"."account_id", "u"."name", "u"."created_at", "u"."updated_at",
	             "a"."name", "a"."created_at", "a"."updated_at"
	        FROM "token" "t"
	       INNER JOIN "user" "u" ON "t"."user_id" = "u"."id"
	       INNER JOIN "account" "a" ON "u"."account_id" = "a"."id"
	       WHERE "t"."id" = $1 AND ("t"."expired_at" IS NULL OR "t"."expired_at" > now())
	         AND "u"."deleted_at" IS NULL
	         AND "a"."deleted_at" IS NULL`
	row := m.conn.QueryRowContext(m.ctx, q, token.ID)
	if err := row.Scan(
		&token.UserID, &token.ExpiredAt, &token.CreatedAt,
		&user.AccountID, &user.Name, &user.CreatedAt, &user.UpdatedAt,
		&account.Name, &account.CreatedAt, &account.UpdatedAt,
	); err != nil {
		return nil, errors.Database(errors.ServerErrorMessage, err, "trying to fetch credentials by the token %q", id)
	}
	user.ID, account.ID = token.UserID, user.AccountID
	token.User, user.Account = &user, &account
	user.Tokens, account.Users = append(user.Tokens, &token), append(account.Users, &user)
	return &token, nil
}
