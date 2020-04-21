package query

import "go.octolab.org/ecosystem/forma/internal/domain"

// CreateSchema TODO issue#173
type CreateSchema struct {
	ID         *domain.ID
	Title      string
	Definition domain.Schema
}

// ReadSchema TODO issue#173
type ReadSchema struct {
	ID domain.ID
}

// UpdateSchema TODO issue#173
type UpdateSchema struct {
	ID         domain.ID
	Title      string
	Definition domain.Schema
}

// DeleteSchema TODO issue#173
type DeleteSchema struct {
	ID          domain.ID
	Permanently bool
}
