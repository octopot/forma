package postgres_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/executor"
	. "github.com/kamilsk/form-api/pkg/storage/executor/internal/postgres"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestLogWriter(t *testing.T) {
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

			e := domain.InputEvent{
				SchemaID: id, InputID: id, TemplateID: &id, Identifier: &id,
				Context: domain.InputContext{}, Code: http.StatusFound, URL: "test",
			}
			l := query.WriteLog{InputEvent: e}

			mock.
				ExpectQuery(`INSERT INTO "event"`).
				WithArgs(
					e.SchemaID, e.InputID, e.TemplateID, e.Identifier,
					[]byte(`{}`), e.Code, e.URL,
				).
				WillReturnRows(
					sqlmock.
						NewRows([]string{"id", "created_at"}).
						AddRow(uint64(1), time.Now()),
				)

			var exec executor.LogWriter = NewEventContext(ctx, conn)
			event, err := exec.Write(l)
			assert.NoError(t, err)
			assert.Equal(t, uint64(1), event.ID)
			assert.NotEmpty(t, event.CreatedAt)
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

			e := domain.InputEvent{
				SchemaID: id, InputID: id, TemplateID: &id, Identifier: &id,
				Context: domain.InputContext{}, Code: http.StatusFound, URL: "test",
			}
			l := query.WriteLog{InputEvent: e}

			mock.
				ExpectQuery(`INSERT INTO "event"`).
				WithArgs(
					e.SchemaID, e.InputID, e.TemplateID, e.Identifier,
					[]byte(`{}`), e.Code, e.URL,
				).
				WillReturnError(errors.Simple("test"))

			var exec executor.LogWriter = NewEventContext(ctx, conn)
			event, err := exec.Write(l)
			assert.Error(t, err)
			assert.Empty(t, event.ID)
			assert.Empty(t, event.CreatedAt)
		})
	})
}
