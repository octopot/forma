package service

import (
	"context"

	"go.octolab.org/ecosystem/forma/internal/domain"
	"go.octolab.org/ecosystem/forma/internal/storage/types"
)

// Storage TODO issue#173
type Storage interface {
	// Schema returns the form schema with provided ID.
	Schema(context.Context, domain.ID) (domain.Schema, error)
	// Template returns the form template with provided ID.
	Template(context.Context, domain.ID) (domain.Template, error)
	// StoreInput stores an user input data.
	StoreInput(context.Context, domain.ID, domain.InputData) (*types.Input, error)
}

// Tracker TODO issue#173
type Tracker interface {
	// LogInput stores an input event.
	LogInput(context.Context, domain.InputEvent) error
}
