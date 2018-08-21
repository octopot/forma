package service

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Storage TODO
type Storage interface {
	// PutData inserts form data and returns their ID.
	PutData(context.Context, domain.ID, domain.InputData) (domain.ID, error)
	// Schema returns the form schema with provided ID.
	Schema(context.Context, domain.ID) (domain.Schema, error)
	// Template returns the form template with provided ID.
	Template(context.Context, domain.ID) (domain.Template, error)
}

// ProtectedStorage TODO
type ProtectedStorage interface {
	Storage
}
