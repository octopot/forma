package query

import "go.octolab.org/ecosystem/forma/internal/domain"

// CreateTemplate TODO issue#173
type CreateTemplate struct {
	ID         *domain.ID
	Title      string
	Definition domain.Template
}

// ReadTemplate TODO issue#173
type ReadTemplate struct {
	ID domain.ID
}

// UpdateTemplate TODO issue#173
type UpdateTemplate struct {
	ID         domain.ID
	Title      string
	Definition domain.Template
}

// DeleteTemplate TODO issue#173
type DeleteTemplate struct {
	ID          domain.ID
	Permanently bool
}
