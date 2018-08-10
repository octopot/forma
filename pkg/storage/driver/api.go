package driver

import "github.com/kamilsk/form-api/pkg/storage"

// CreateSchema TODO
type CreateSchema struct {
	Language   string
	Title      string
	Definition string
}

// ReadSchema TODO
type ReadSchema struct {
	ID string
}

// UpdateSchema TODO
type UpdateSchema struct {
	ID         string
	Language   string
	Title      string
	Definition string
}

// DeleteSchema TODO
type DeleteSchema struct {
	ID          string
	Permanently bool
}

// SchemaEditor TODO
type SchemaEditor interface {
	Create(*storage.Token, CreateSchema) (storage.Schema, error)
	Read(*storage.Token, ReadSchema) (storage.Schema, error)
	Update(*storage.Token, UpdateSchema) (storage.Schema, error)
	Delete(*storage.Token, DeleteSchema) (storage.Schema, error)
}

// CreateTemplate TODO
type CreateTemplate struct {
	Title      string
	Definition string
}

// ReadTemplate TODO
type ReadTemplate struct {
	ID string
}

// UpdateTemplate TODO
type UpdateTemplate struct {
	ID         string
	Title      string
	Definition string
}

// DeleteTemplate TODO
type DeleteTemplate struct {
	ID          string
	Permanently bool
}

// TemplateEditor TODO
type TemplateEditor interface {
	Create(*storage.Token, CreateTemplate) (storage.Template, error)
	Read(*storage.Token, ReadTemplate) (storage.Template, error)
	Update(*storage.Token, UpdateTemplate) (storage.Template, error)
	Delete(*storage.Token, DeleteTemplate) (storage.Template, error)
}
