package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/kamilsk/form-api/pkg/storage/executor"
	"github.com/kamilsk/form-api/pkg/storage/executor/internal/postgres"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNewInputContext(t *testing.T) {
	token := Token()
	t.Run("read by ID", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		conn, err := db.Conn(ctx)
		assert.NoError(t, err)
		defer conn.Close()

		mock.
			ExpectQuery(`SELECT "(?:.+)" FROM "input"`).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(
				sqlmock.
					NewRows([]string{"schema_id", "data", "created_at"}).
					AddRow("10000000-2000-4000-8000-160000000004", "{}", time.Now()),
			)

		var exec executor.InputReader = postgres.NewInputContext(ctx, conn)
		input, err := exec.ReadByID(token, "10000000-2000-4000-8000-160000000000")
		assert.NoError(t, err)
		assert.Equal(t, "10000000-2000-4000-8000-160000000004", input.SchemaID)
	})
	t.Run("read by filter", func(t *testing.T) {
		// TODO
	})
}
