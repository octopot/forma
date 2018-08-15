package executor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kamilsk/form-api/pkg/storage/executor/internal/postgres"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

const (
	postgresDialect = "postgres"
	mysqlDialect    = "mysql"
)

// New TODO
func New(dialect string) *Executor {
	exec := &Executor{dialect: dialect}
	switch exec.dialect {
	case postgresDialect:
		exec.factory.NewInputReader = func(conn *sql.Conn, ctx context.Context) InputReader {
			return postgres.NewInputContext(conn, ctx)
		}
		exec.factory.NewInputWriter = func(conn *sql.Conn, ctx context.Context) InputWriter {
			return postgres.NewInputContext(conn, ctx)
		}
		exec.factory.NewSchemaEditor = func(conn *sql.Conn, ctx context.Context) SchemaEditor {
			return postgres.NewSchemaContext(conn, ctx)
		}
		exec.factory.NewSchemaReader = func(conn *sql.Conn, ctx context.Context) SchemaReader {
			return postgres.NewSchemaContext(conn, ctx)
		}
		exec.factory.NewTemplateEditor = func(conn *sql.Conn, ctx context.Context) TemplateEditor {
			return postgres.NewTemplateContext(conn, ctx)
		}
		exec.factory.NewTemplateReader = func(conn *sql.Conn, ctx context.Context) TemplateReader {
			return postgres.NewTemplateContext(conn, ctx)
		}
		exec.factory.NewUserManager = func(conn *sql.Conn, ctx context.Context) UserManager {
			return postgres.NewUserContext(conn, ctx)
		}
	case mysqlDialect:
		fallthrough
	default:
		panic(fmt.Sprintf("not supported dialect %q is provided", exec.dialect))
	}
	return exec
}

// InputReader TODO
type InputReader interface {
	ReadByID(*query.Token, string) (query.Input, error)
	ReadByFilter(*query.Token, query.InputFilter) ([]query.Input, error)
}

// InputWriter TODO
type InputWriter interface {
	Write(query.WriteInput) (query.Input, error)
}

// SchemaEditor TODO
type SchemaEditor interface {
	Create(*query.Token, query.CreateSchema) (query.Schema, error)
	Read(*query.Token, query.ReadSchema) (query.Schema, error)
	Update(*query.Token, query.UpdateSchema) (query.Schema, error)
	Delete(*query.Token, query.DeleteSchema) (query.Schema, error)
}

// SchemaReader TODO
type SchemaReader interface {
	ReadByID(string) (query.Schema, error)
}

// TemplateEditor TODO
type TemplateEditor interface {
	Create(*query.Token, query.CreateTemplate) (query.Template, error)
	Read(*query.Token, query.ReadTemplate) (query.Template, error)
	Update(*query.Token, query.UpdateTemplate) (query.Template, error)
	Delete(*query.Token, query.DeleteTemplate) (query.Template, error)
}

// TemplateReader TODO
type TemplateReader interface {
	ReadByID(string) (query.Template, error)
}

// UserManager TODO
type UserManager interface {
	Token(string) (*query.Token, error)
}

// Executor TODO
type Executor struct {
	dialect string
	factory struct {
		NewInputReader    func(*sql.Conn, context.Context) InputReader
		NewInputWriter    func(*sql.Conn, context.Context) InputWriter
		NewSchemaEditor   func(*sql.Conn, context.Context) SchemaEditor
		NewSchemaReader   func(*sql.Conn, context.Context) SchemaReader
		NewTemplateEditor func(*sql.Conn, context.Context) TemplateEditor
		NewTemplateReader func(*sql.Conn, context.Context) TemplateReader
		NewUserManager    func(*sql.Conn, context.Context) UserManager
	}
}

// Dialect TODO
func (e *Executor) Dialect() string {
	return e.dialect
}

// InputReader TODO
func (e *Executor) InputReader(conn *sql.Conn, ctx context.Context) InputReader {
	return e.factory.NewInputReader(conn, ctx)
}

// InputWriter TODO
func (e *Executor) InputWriter(conn *sql.Conn, ctx context.Context) InputWriter {
	return e.factory.NewInputWriter(conn, ctx)
}

// SchemaEditor TODO
func (e *Executor) SchemaEditor(conn *sql.Conn, ctx context.Context) SchemaEditor {
	return e.factory.NewSchemaEditor(conn, ctx)
}

// SchemaReader TODO
func (e *Executor) SchemaReader(conn *sql.Conn, ctx context.Context) SchemaReader {
	return e.factory.NewSchemaReader(conn, ctx)
}

// TemplateEditor TODO
func (e *Executor) TemplateEditor(conn *sql.Conn, ctx context.Context) TemplateEditor {
	return e.factory.NewTemplateEditor(conn, ctx)
}

// TemplateReader TODO
func (e *Executor) TemplateReader(conn *sql.Conn, ctx context.Context) TemplateReader {
	return e.factory.NewTemplateReader(conn, ctx)
}

// UserManager TODO
func (e *Executor) UserManager(conn *sql.Conn, ctx context.Context) UserManager {
	return e.factory.NewUserManager(conn, ctx)
}
