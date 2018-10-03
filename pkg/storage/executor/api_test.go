package executor_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/kamilsk/form-api/pkg/storage/executor"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type contract interface {
		Dialect() string

		InputReader(context.Context, *sql.Conn) executor.InputReader
		InputWriter(context.Context, *sql.Conn) executor.InputWriter
		LogWriter(context.Context, *sql.Conn) executor.LogWriter
		SchemaEditor(context.Context, *sql.Conn) executor.SchemaEditor
		TemplateEditor(context.Context, *sql.Conn) executor.TemplateEditor
		UserManager(context.Context, *sql.Conn) executor.UserManager

		// Deprecated TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
		SchemaReader(context.Context, *sql.Conn) executor.SchemaReader
		// Deprecated TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
		TemplateReader(context.Context, *sql.Conn) executor.TemplateReader
	}
	t.Run("PostgreSQL", func(t *testing.T) {
		assert.NotPanics(t, func() {
			dialect, ctx := "postgres", context.Background()
			var exec contract = executor.New(dialect)
			assert.Equal(t, dialect, exec.Dialect())

			assert.NotNil(t, exec.InputReader(ctx, nil))
			assert.NotNil(t, exec.InputWriter(ctx, nil))
			assert.NotNil(t, exec.LogWriter(ctx, nil))
			assert.NotNil(t, exec.SchemaEditor(ctx, nil))
			assert.NotNil(t, exec.TemplateEditor(ctx, nil))
			assert.NotNil(t, exec.UserManager(ctx, nil))

			assert.NotNil(t, exec.SchemaReader(ctx, nil))
			assert.NotNil(t, exec.TemplateReader(ctx, nil))
		})
	})
	t.Run("MySQL", func(t *testing.T) {
		assert.Panics(t, func() {
			dialect := "mysql"
			var _ contract = executor.New(dialect)
		})
	})
}
