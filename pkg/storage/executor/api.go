package executor

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kamilsk/form-api/pkg/domain"
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
		exec.factory.NewInputReader = func(ctx context.Context, conn *sql.Conn) InputReader {
			return postgres.NewInputContext(ctx, conn)
		}
		exec.factory.NewInputWriter = func(ctx context.Context, conn *sql.Conn) InputWriter {
			return postgres.NewInputContext(ctx, conn)
		}
		exec.factory.NewLogWriter = func(ctx context.Context, conn *sql.Conn) LogWriter {
			return postgres.NewLogContext(ctx, conn)
		}
		exec.factory.NewSchemaEditor = func(ctx context.Context, conn *sql.Conn) SchemaEditor {
			return postgres.NewSchemaContext(ctx, conn)
		}
		exec.factory.NewSchemaReader = func(ctx context.Context, conn *sql.Conn) SchemaReader {
			return postgres.NewSchemaContext(ctx, conn)
		}
		exec.factory.NewTemplateEditor = func(ctx context.Context, conn *sql.Conn) TemplateEditor {
			return postgres.NewTemplateContext(ctx, conn)
		}
		exec.factory.NewTemplateReader = func(ctx context.Context, conn *sql.Conn) TemplateReader {
			return postgres.NewTemplateContext(ctx, conn)
		}
		exec.factory.NewUserManager = func(ctx context.Context, conn *sql.Conn) UserManager {
			return postgres.NewUserContext(ctx, conn)
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
	ReadByID(*query.Token, domain.ID) (query.Input, error)
	ReadByFilter(*query.Token, query.InputFilter) ([]query.Input, error)
}

// InputWriter TODO
type InputWriter interface {
	Write(query.WriteInput) (query.Input, error)
}

// LogWriter TODO
type LogWriter interface {
	Write(query.WriteLog) (query.Log, error)
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
	ReadByID(domain.ID) (query.Schema, error)
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
	ReadByID(domain.ID) (query.Template, error)
}

// UserManager TODO
type UserManager interface {
	Token(domain.ID) (*query.Token, error)
}

// Executor TODO
type Executor struct {
	dialect string
	factory struct {
		NewInputReader    func(context.Context, *sql.Conn) InputReader
		NewInputWriter    func(context.Context, *sql.Conn) InputWriter
		NewLogWriter      func(context.Context, *sql.Conn) LogWriter
		NewSchemaEditor   func(context.Context, *sql.Conn) SchemaEditor
		NewSchemaReader   func(context.Context, *sql.Conn) SchemaReader
		NewTemplateEditor func(context.Context, *sql.Conn) TemplateEditor
		NewTemplateReader func(context.Context, *sql.Conn) TemplateReader
		NewUserManager    func(context.Context, *sql.Conn) UserManager
	}
}

// Dialect TODO
func (e *Executor) Dialect() string {
	return e.dialect
}

// InputReader TODO
func (e *Executor) InputReader(ctx context.Context, conn *sql.Conn) InputReader {
	return e.factory.NewInputReader(ctx, conn)
}

// InputWriter TODO
func (e *Executor) InputWriter(ctx context.Context, conn *sql.Conn) InputWriter {
	return e.factory.NewInputWriter(ctx, conn)
}

// LogWriter TODO
func (e *Executor) LogWriter(ctx context.Context, conn *sql.Conn) LogWriter {
	return e.factory.NewLogWriter(ctx, conn)
}

// SchemaEditor TODO
func (e *Executor) SchemaEditor(ctx context.Context, conn *sql.Conn) SchemaEditor {
	return e.factory.NewSchemaEditor(ctx, conn)
}

// SchemaReader TODO
func (e *Executor) SchemaReader(ctx context.Context, conn *sql.Conn) SchemaReader {
	return e.factory.NewSchemaReader(ctx, conn)
}

// TemplateEditor TODO
func (e *Executor) TemplateEditor(ctx context.Context, conn *sql.Conn) TemplateEditor {
	return e.factory.NewTemplateEditor(ctx, conn)
}

// TemplateReader TODO
func (e *Executor) TemplateReader(ctx context.Context, conn *sql.Conn) TemplateReader {
	return e.factory.NewTemplateReader(ctx, conn)
}

// UserManager TODO
func (e *Executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return e.factory.NewUserManager(ctx, conn)
}
