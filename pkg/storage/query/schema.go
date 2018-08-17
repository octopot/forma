package query

import "github.com/kamilsk/form-api/pkg/domain"

// CreateSchema TODO
type CreateSchema struct {
	Language   string
	Title      string
	Definition string
}

// ReadSchema TODO
type ReadSchema struct {
	ID domain.ID
}

// UpdateSchema TODO
type UpdateSchema struct {
	ID         domain.ID
	Language   string
	Title      string
	Definition string
}

// DeleteSchema TODO
type DeleteSchema struct {
	ID          domain.ID
	Permanently bool
}
