package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage"
	"github.com/kamilsk/form-api/pkg/storage/driver"
)

// NewUserContext TODO
func NewUserContext(conn *sql.Conn, ctx context.Context) (driver.UserManager, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	return manager{conn, ctx}, func() {
		cancel()
		_ = conn.Close()
	}
}

type manager struct {
	conn *sql.Conn
	ctx  context.Context
}

// Token TODO
func (m manager) Token(id string) (*storage.Token, error) {
	var (
		token   = storage.Token{ID: id}
		user    = storage.User{}
		account = storage.Account{}
	)
	query := `SELECT "t"."user_id", "t"."expired_at", "t"."created_at",
	                 "u"."account_id", "u"."name", "u"."created_at", "u"."updated_at",
	                 "a"."name", "a"."created_at", "a"."updated_at"
	            FROM "token" "t"
	           INNER JOIN "user" "u" ON "t"."user_id" = "u"."id"
	           INNER JOIN "account" "a" ON "u"."account_id" = "a"."id"
	           WHERE "t"."id" = $1 AND ("t"."expired_at" IS NULL OR "t"."expired_at" > now())
	             AND "u"."deleted_at" IS NULL
	             AND "a"."deleted_at" IS NULL`
	row := m.conn.QueryRowContext(m.ctx, query, token.ID)
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
