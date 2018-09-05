package service

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// Storage TODO
type Storage interface {
	// Schema returns the form schema with provided ID.
	Schema(context.Context, domain.ID) (domain.Schema, error)
	// Template returns the form template with provided ID.
	Template(context.Context, domain.ID) (domain.Template, error)
}

// InputHandler TODO
type InputHandler interface {
	// HandleInput TODO
	HandleInput(context.Context, domain.ID, domain.InputData) (*types.Input, error)
	// LogRequest TODO
	LogRequest(context.Context, *types.Input, domain.InputContext) error
}
