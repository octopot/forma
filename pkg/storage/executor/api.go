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
func New(dialect string) *executor {
	exec := &executor{dialect: dialect}
	switch exec.dialect {
	case postgresDialect:
		exec.factory.NewInputReader = func(conn *sql.Conn, ctx context.Context) InputReader {
			return postgres.NewInputContext(conn, ctx)
		}
		exec.factory.NewSchemaEditor = func(conn *sql.Conn, ctx context.Context) SchemaEditor {
			return postgres.NewSchemaContext(conn, ctx)
		}
		exec.factory.NewTemplateEditor = func(conn *sql.Conn, ctx context.Context) TemplateEditor {
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

// TemplateEditor TODO
type TemplateEditor interface {
	Create(*query.Token, query.CreateTemplate) (query.Template, error)
	Read(*query.Token, query.ReadTemplate) (query.Template, error)
	Update(*query.Token, query.UpdateTemplate) (query.Template, error)
	Delete(*query.Token, query.DeleteTemplate) (query.Template, error)
}

// UserManager TODO
type UserManager interface {
	Token(string) (*query.Token, error)
}

type executor struct {
	dialect string
	factory struct {
		NewInputReader    func(*sql.Conn, context.Context) InputReader
		NewSchemaEditor   func(*sql.Conn, context.Context) SchemaEditor
		NewTemplateEditor func(*sql.Conn, context.Context) TemplateEditor
		NewUserManager    func(*sql.Conn, context.Context) UserManager
	}
}

// Dialect TODO
func (e *executor) Dialect() string {
	return e.dialect
}

// InputReader TODO
func (e *executor) InputReader(conn *sql.Conn, ctx context.Context) InputReader {
	return e.factory.NewInputReader(conn, ctx)
}

// SchemaEditor TODO
func (e *executor) SchemaEditor(conn *sql.Conn, ctx context.Context) SchemaEditor {
	return e.factory.NewSchemaEditor(conn, ctx)
}

// TemplateEditor TODO
func (e *executor) TemplateEditor(conn *sql.Conn, ctx context.Context) TemplateEditor {
	return e.factory.NewTemplateEditor(conn, ctx)
}

// UserManager TODO
func (e *executor) UserManager(conn *sql.Conn, ctx context.Context) UserManager {
	return e.factory.NewUserManager(conn, ctx)
}
