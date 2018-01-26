package service

import "github.com/kamilsk/form-api/domain"

// Storage defines the behavior of Data Access Object.
type Storage interface {
	// AddData inserts form data and returns their ID.
	AddData(uuid domain.UUID, values map[string][]string) (int64, error)
	// Schema returns the form schema with provided UUID.
	Schema(uuid domain.UUID) (domain.Schema, error)
}
