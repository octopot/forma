package service

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// Storage TODO issue#173
type Storage interface {
	// Schema returns the form schema with provided ID.
	Schema(context.Context, domain.ID) (domain.Schema, error)
	// Template returns the form template with provided ID.
	Template(context.Context, domain.ID) (domain.Template, error)
}

// InputHandler TODO issue#173
type InputHandler interface {
	// HandleInput TODO issue#173
	HandleInput(context.Context, domain.ID, domain.InputData) (*types.Input, error)
	// LogRequest TODO issue#173
	LogRequest(context.Context, *types.Input, domain.InputContext) error
}
