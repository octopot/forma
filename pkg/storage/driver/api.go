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
