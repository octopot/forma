package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"go.octolab.org/ecosystem/forma/internal/domain"
	"go.octolab.org/ecosystem/forma/internal/errors"
	"go.octolab.org/ecosystem/forma/internal/storage/executor"
	. "go.octolab.org/ecosystem/forma/internal/storage/executor/internal/postgres"
	"go.octolab.org/ecosystem/forma/internal/storage/query"
)

func TestInputReader(t *testing.T) {
	token, id := DemoToken(), domain.ID("10000000-2000-4000-8000-160000000000")
	t.Run("read by ID", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer func() { _ = conn.Close() }()

			mock.
				ExpectQuery(`SELECT "(?:.+)" FROM "input"`).
				WithArgs(id, token.User.AccountID).
				WillReturnRows(
					sqlmock.
						NewRows([]string{"schema_id", "data", "created_at"}).
						AddRow(id, `{"input":["test"]}`, time.Now()),
				)

			var exec executor.InputReader = NewInputContext(ctx, conn)
			input, err := exec.ReadByID(token, id)
			assert.NoError(t, err)
			assert.Equal(t, id, input.SchemaID)
			assert.Equal(t, domain.InputData{"input": {"test"}}, input.Data)
			assert.NotEmpty(t, input.CreatedAt)
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
				ExpectQuery(`SELECT "(?:.+)" FROM "input"`).
				WithArgs(id, token.User.AccountID).
				WillReturnError(errors.Simple("test"))

			var exec executor.InputReader = NewInputContext(ctx, conn)
			input, err := exec.ReadByID(token, id)
			assert.Error(t, err)
			assert.Empty(t, input.SchemaID)
			assert.Empty(t, input.Data)
			assert.Empty(t, input.CreatedAt)
		})
		t.Run("serialization error", func(t *testing.T) {
			// TODO issue#126
		})
	})
	t.Run("read by filter", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer func() { _ = conn.Close() }()

			now := time.Now()

			tests := []struct {
				name   string
				mocker func(sqlmock.Sqlmock)
				filter query.InputFilter
			}{
				{
					name: "by schema ID",
					mocker: func(mock sqlmock.Sqlmock) {
						mock.
							ExpectQuery(`SELECT "(?:.+)" FROM "input"`).
							WithArgs(id, token.User.AccountID).
							WillReturnRows(
								sqlmock.
									NewRows([]string{"id", "data", "created_at"}).
									AddRow(id, `{"input":["test"]}`, now),
							)
					},
					filter: query.InputFilter{SchemaID: id},
				},
				{
					name: `by schema ID and "from" date`,
					mocker: func(mock sqlmock.Sqlmock) {
						mock.
							ExpectQuery(`SELECT "(?:.+)" FROM "input"`).
							WithArgs(id, token.User.AccountID, sqlmock.AnyArg()).
							WillReturnRows(
								sqlmock.
									NewRows([]string{"id", "data", "created_at"}).
									AddRow(id, `{"input":["test"]}`, now),
							)
					},
					filter: query.InputFilter{SchemaID: id, From: &now},
				},
				{
					name: `by schema ID and "to" date`,
					mocker: func(mock sqlmock.Sqlmock) {
						mock.
							ExpectQuery(`SELECT "(?:.+)" FROM "input"`).
							WithArgs(id, token.User.AccountID, sqlmock.AnyArg()).
							WillReturnRows(
								sqlmock.
									NewRows([]string{"id", "data", "created_at"}).
									AddRow(id, `{"input":["test"]}`, now),
							)
					},
					filter: query.InputFilter{SchemaID: id, To: &now},
				},
				{
					name: `by schema ID and "from" and "to" dates`,
					mocker: func(mock sqlmock.Sqlmock) {
						mock.
							ExpectQuery(`SELECT "(?:.+)" FROM "input"`).
							WithArgs(id, token.User.AccountID, sqlmock.AnyArg(), sqlmock.AnyArg()).
							WillReturnRows(
								sqlmock.
									NewRows([]string{"id", "data", "created_at"}).
									AddRow(id, `{"input":["test"]}`, now),
							)
					},
					filter: query.InputFilter{SchemaID: id, From: &now, To: &now},
				},
			}

			var exec executor.InputReader = NewInputContext(ctx, conn)
			for _, test := range tests {
				test.mocker(mock)
				inputs, readErr := exec.ReadByFilter(token, test.filter)
				assert.NoError(t, readErr, test.name)
				assert.Len(t, inputs, 1, test.name)
			}
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
				ExpectQuery(`SELECT "(?:.+)" FROM "input"`).
				WithArgs(id, token.User.AccountID).
				WillReturnError(errors.Simple("test"))

			var exec executor.InputReader = NewInputContext(ctx, conn)
			inputs, err := exec.ReadByFilter(token, query.InputFilter{SchemaID: id})
			assert.Error(t, err)
			assert.Nil(t, inputs)
		})
		t.Run("database scan error", func(t *testing.T) {
			// TODO issue#126
		})
		t.Run("serialization error", func(t *testing.T) {
			// TODO issue#126
		})
	})
}

func TestInputWriter(t *testing.T) {
	id := domain.ID("10000000-2000-4000-8000-160000000000")
	t.Run("write", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer func() { _ = conn.Close() }()

			mock.
				ExpectQuery(`INSERT INTO "input"`).
				WithArgs(id, []byte(`{"input":["test"]}`)).
				WillReturnRows(
					sqlmock.
						NewRows([]string{"id", "created_at"}).
						AddRow(id, time.Now()),
				)

			var exec executor.InputWriter = NewInputContext(ctx, conn)
			input, err := exec.Write(query.WriteInput{
				SchemaID:     id,
				VerifiedData: map[string][]string{"input": {"test"}},
			})
			assert.NoError(t, err)
			assert.Equal(t, id, input.ID)
			assert.NotEmpty(t, input.CreatedAt)
		})
		t.Run("serialization error", func(t *testing.T) {
			// TODO issue#126
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
				ExpectQuery(`INSERT INTO "input"`).
				WithArgs(id, []byte(`{"input":["test"]}`)).
				WillReturnError(errors.Simple("test"))

			var exec executor.InputWriter = NewInputContext(ctx, conn)
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
