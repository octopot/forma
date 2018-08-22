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
		SchemaReader(context.Context, *sql.Conn) executor.SchemaReader
		TemplateEditor(context.Context, *sql.Conn) executor.TemplateEditor
		TemplateReader(context.Context, *sql.Conn) executor.TemplateReader
		UserManager(context.Context, *sql.Conn) executor.UserManager
	}
	t.Run("PostgreSQL", func(t *testing.T) {
		assert.NotPanics(t, func() {
			var exec contract = executor.New("postgres")
			assert.NotEmpty(t, exec.Dialect())
			assert.NotNil(t, exec.InputReader(nil, nil))
			assert.NotNil(t, exec.InputWriter(nil, nil))
			assert.NotNil(t, exec.LogWriter(nil, nil))
			assert.NotNil(t, exec.SchemaEditor(nil, nil))
			assert.NotNil(t, exec.SchemaReader(nil, nil))
			assert.NotNil(t, exec.TemplateEditor(nil, nil))
			assert.NotNil(t, exec.TemplateReader(nil, nil))
			assert.NotNil(t, exec.UserManager(nil, nil))
		})
	})
	t.Run("MySQL", func(t *testing.T) {
		assert.Panics(t, func() {
			var _ contract = executor.New("mysql")
		})
	})
}
