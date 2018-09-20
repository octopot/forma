package executor

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/executor/internal/postgres"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

const (
	postgresDialect = "postgres"
	mysqlDialect    = "mysql"
)

// New TODO issue#173
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
		exec.factory.NewTemplateEditor = func(ctx context.Context, conn *sql.Conn) TemplateEditor {
			return postgres.NewTemplateContext(ctx, conn)
		}
		exec.factory.NewUserManager = func(ctx context.Context, conn *sql.Conn) UserManager {
			return postgres.NewUserContext(ctx, conn)
		}
		// Deprecated TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
		exec.factory.NewSchemaReader = func(ctx context.Context, conn *sql.Conn) SchemaReader {
			return postgres.NewSchemaContext(ctx, conn)
		}
		// Deprecated TODO issue#version3.0 use TemplateEditor and gRPC gateway instead
		exec.factory.NewTemplateReader = func(ctx context.Context, conn *sql.Conn) TemplateReader {
			return postgres.NewTemplateContext(ctx, conn)
		}
	case mysqlDialect:
		fallthrough
	default:
		panic(fmt.Sprintf("not supported dialect %q is provided", exec.dialect))
	}
	return exec
}

// InputReader TODO issue#173
type InputReader interface {
	ReadByID(*types.Token, domain.ID) (types.Input, error)
	ReadByFilter(*types.Token, query.InputFilter) ([]types.Input, error)
}

// InputWriter TODO issue#173
type InputWriter interface {
	Write(query.WriteInput) (types.Input, error)
}

// LogWriter TODO issue#173
type LogWriter interface {
	Write(query.WriteLog) (types.Log, error)
}

// SchemaEditor TODO issue#173
type SchemaEditor interface {
	Create(*types.Token, query.CreateSchema) (types.Schema, error)
	Read(*types.Token, query.ReadSchema) (types.Schema, error)
	Update(*types.Token, query.UpdateSchema) (types.Schema, error)
	Delete(*types.Token, query.DeleteSchema) (types.Schema, error)
}

// SchemaReader TODO issue#173
// Deprecated TODO issue#version3.0 use TemplateEditor and gRPC gateway instead
type SchemaReader interface {
	ReadByID(domain.ID) (types.Schema, error)
}

// TemplateEditor TODO issue#173
type TemplateEditor interface {
	Create(*types.Token, query.CreateTemplate) (types.Template, error)
	Read(*types.Token, query.ReadTemplate) (types.Template, error)
	Update(*types.Token, query.UpdateTemplate) (types.Template, error)
	Delete(*types.Token, query.DeleteTemplate) (types.Template, error)
}

// TemplateReader TODO issue#173
// Deprecated TODO issue#version3.0 use TemplateEditor and gRPC gateway instead
type TemplateReader interface {
	ReadByID(domain.ID) (types.Template, error)
}

// UserManager TODO issue#173
type UserManager interface {
	Token(domain.ID) (*types.Token, error)
}

// Executor TODO issue#173
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

// Dialect TODO issue#173
func (e *Executor) Dialect() string {
	return e.dialect
}

// InputReader TODO issue#173
func (e *Executor) InputReader(ctx context.Context, conn *sql.Conn) InputReader {
	return e.factory.NewInputReader(ctx, conn)
}

// InputWriter TODO issue#173
func (e *Executor) InputWriter(ctx context.Context, conn *sql.Conn) InputWriter {
	return e.factory.NewInputWriter(ctx, conn)
}

// LogWriter TODO issue#173
func (e *Executor) LogWriter(ctx context.Context, conn *sql.Conn) LogWriter {
	return e.factory.NewLogWriter(ctx, conn)
}

// SchemaEditor TODO issue#173
func (e *Executor) SchemaEditor(ctx context.Context, conn *sql.Conn) SchemaEditor {
	return e.factory.NewSchemaEditor(ctx, conn)
}

// TemplateEditor TODO issue#173
func (e *Executor) TemplateEditor(ctx context.Context, conn *sql.Conn) TemplateEditor {
	return e.factory.NewTemplateEditor(ctx, conn)
}

// UserManager TODO issue#173
func (e *Executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return e.factory.NewUserManager(ctx, conn)
}

// SchemaReader TODO issue#173
// Deprecated TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
func (e *Executor) SchemaReader(ctx context.Context, conn *sql.Conn) SchemaReader {
	return e.factory.NewSchemaReader(ctx, conn)
}

// TemplateReader TODO issue#173
// Deprecated TODO issue#version3.0 use TemplateEditor and gRPC gateway instead
func (e *Executor) TemplateReader(ctx context.Context, conn *sql.Conn) TemplateReader {
	return e.factory.NewTemplateReader(ctx, conn)
}
