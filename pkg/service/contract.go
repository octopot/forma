package service

import "github.com/kamilsk/form-api/pkg/domain"

// Storage TODO
type Storage interface {
	// AddData inserts form data and returns their ID.
	AddData(domain.UUID, map[string][]string) (int64, error)
	// Schema returns the form schema with provided UUID.
	Schema(domain.UUID) (domain.Schema, error)
	// Template returns the form template with provided UUID.
	// Template(domain.UUID) (domain.Template, error)
}

// ProtectedStorage TODO
type ProtectedStorage interface {
	Storage
}
