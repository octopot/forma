package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/executor"
	"github.com/kamilsk/form-api/pkg/storage/executor/internal/postgres"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestInputReader(t *testing.T) {
	token, id := Token(), "10000000-2000-4000-8000-160000000000"
	t.Run("read by ID", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
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
						AddRow(id, `{"input":["test"]}`, time.Now()),
				)

			var exec executor.InputReader = postgres.NewInputContext(ctx, conn)
			input, err := exec.ReadByID(token, id)
			assert.NoError(t, err)
			assert.Equal(t, id, input.SchemaID)
			assert.Equal(t, []byte(`{"input":["test"]}`), input.Data)
			assert.NotEmpty(t, input.CreatedAt)
		})
		t.Run("database error", func(t *testing.T) {
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
				WillReturnError(errors.Simple("test"))

			var exec executor.InputReader = postgres.NewInputContext(ctx, conn)
			input, err := exec.ReadByID(token, id)
			assert.Error(t, err)
			assert.Empty(t, input.SchemaID)
			assert.Empty(t, input.Data)
			assert.Empty(t, input.CreatedAt)
		})
	})
	t.Run("read by filter", func(t *testing.T) {
		// TODO
	})
}

func TestInputWriter(t *testing.T) {
	var id = "10000000-2000-4000-8000-160000000000"
	t.Run("write", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer conn.Close()

			mock.
				ExpectQuery(`INSERT INTO "input"`).
				WithArgs(id, []byte(`{"input":["test"]}`)).
				WillReturnRows(
					sqlmock.
						NewRows([]string{"id", "created_at"}).
						AddRow(id, time.Now()),
				)

			var exec executor.InputWriter = postgres.NewInputContext(ctx, conn)
			input, err := exec.Write(query.WriteInput{
				SchemaID:     id,
				VerifiedData: map[string][]string{"input": {"test"}},
			})
			assert.NoError(t, err)
			assert.Equal(t, id, input.ID)
			assert.NotEmpty(t, input.CreatedAt)
		})
		t.Run("serialization error", func(t *testing.T) {
			// TODO
		})
		t.Run("database error", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer conn.Close()

			mock.
				ExpectQuery(`INSERT INTO "input"`).
				WithArgs(id, []byte(`{"input":["test"]}`)).
				WillReturnError(errors.Simple("test"))

			var exec executor.InputWriter = postgres.NewInputContext(ctx, conn)
			input, err := exec.Write(query.WriteInput{
				SchemaID:     id,
				VerifiedData: map[string][]string{"input": {"test"}},
			})
			assert.Error(t, err)
			assert.Empty(t, input.ID)
			assert.Empty(t, input.CreatedAt)
		})
	})
}
