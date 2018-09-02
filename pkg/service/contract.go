package service

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// Storage TODO
type Storage interface {
	// Schema returns the form schema with provided ID.
	Schema(context.Context, domain.ID) (domain.Schema, error)
	// Template returns the form template with provided ID.
	Template(context.Context, domain.ID) (domain.Template, error)
}

// ProtectedStorage TODO
type ProtectedStorage interface {
	Storage
}

// InputHandler TODO
type InputHandler interface {
	// HandleInput TODO
	HandleInput(context.Context, domain.ID, domain.InputData) (*query.Input, error)
	// LogRequest TODO
	LogRequest(context.Context, *query.Input, domain.InputContext) error
}
