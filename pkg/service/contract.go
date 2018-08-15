package service

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Storage TODO
type Storage interface {
	// AddData inserts form data and returns their ID.
	AddData(context.Context, domain.UUID, domain.InputData) (domain.UUID, error)
	// Schema returns the form schema with provided UUID.
	Schema(context.Context, domain.UUID) (domain.Schema, error)
	// Template returns the form template with provided UUID.
	Template(context.Context, domain.UUID) (domain.Template, error)
}

// ProtectedStorage TODO
type ProtectedStorage interface {
	Storage
}
