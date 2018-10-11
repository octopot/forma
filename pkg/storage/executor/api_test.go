package executor_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/form-api/pkg/storage/executor"
)

func TestNew(t *testing.T) {
	type contract interface {
		Dialect() string

		InputReader(context.Context, *sql.Conn) InputReader
		InputWriter(context.Context, *sql.Conn) InputWriter
		LogWriter(context.Context, *sql.Conn) LogWriter
		SchemaEditor(context.Context, *sql.Conn) SchemaEditor
		TemplateEditor(context.Context, *sql.Conn) TemplateEditor
		UserManager(context.Context, *sql.Conn) UserManager

		// Deprecated TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
		SchemaReader(context.Context, *sql.Conn) SchemaReader
		// Deprecated TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
		TemplateReader(context.Context, *sql.Conn) TemplateReader
	}
	t.Run("PostgreSQL", func(t *testing.T) {
		assert.NotPanics(t, func() {
			dialect, ctx := "postgres", context.Background()
			var exec contract = New(dialect)
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
			var _ contract = New(dialect)
		})
	})
}
