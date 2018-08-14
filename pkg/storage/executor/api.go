package executor

import "github.com/kamilsk/form-api/pkg/storage/query"

const (
	postgresDialect = "postgres"
	mysqlDialect    = "mysql"
)

// UserManager TODO
type UserManager interface {
	Token(string) (*query.Token, error)
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

// InputReader TODO
type InputReader interface {
	ReadByID(*query.Token, string) (query.Input, error)
	ReadByFilter(*query.Token, query.InputFilter) ([]query.Input, error)
}
