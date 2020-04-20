package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/executor"
	. "github.com/kamilsk/form-api/pkg/storage/executor/internal/postgres"
)

func TestUserManager(t *testing.T) {
	id := domain.ID("10000000-2000-4000-8000-160000000000")
	t.Run("token", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer func() { _ = conn.Close() }()

			mock.
				ExpectQuery(`SELECT "(?:.+)" FROM "token"`).
				WithArgs(id).
				WillReturnRows(
					sqlmock.
						NewRows([]string{
							"user_id", "expired_at", "created_at",
							"account_id", "name", "created_at", "updated_at",
							"name", "created_at", "updated_at",
						}).
						AddRow(
							id, time.Now(), time.Now(),
							id, "test", time.Now(), time.Now(),
							"test", time.Now(), time.Now(),
						),
				)

			var exec executor.UserManager = NewUserContext(ctx, conn)
			token, err := exec.Token(id)
			assert.NoError(t, err)
			assert.NotEmpty(t, token.UserID)
			assert.NotEmpty(t, token.ExpiredAt)
			assert.NotEmpty(t, token.CreatedAt)
			assert.NotEmpty(t, token.User)
			assert.NotEmpty(t, token.User.AccountID)
			assert.NotEmpty(t, token.User.Name)
			assert.NotEmpty(t, token.User.CreatedAt)
			assert.NotEmpty(t, token.User.UpdatedAt)
			assert.NotEmpty(t, token.User.Account)
			assert.NotEmpty(t, token.User.Account.Name)
			assert.NotEmpty(t, token.User.Account.CreatedAt)
			assert.NotEmpty(t, token.User.Account.UpdatedAt)
		})
		t.Run("database error", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer func() { _ = conn.Close() }()

			mock.
				ExpectQuery(`SELECT "(?:.+)" FROM "token"`).
				WithArgs(id).
				WillReturnError(errors.Simple("test"))

			var exec executor.UserManager = NewUserContext(ctx, conn)
			token, err := exec.Token(id)
			assert.Error(t, err)
			assert.Nil(t, token)
		})
	})
}
