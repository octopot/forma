package query

import "github.com/kamilsk/form-api/pkg/domain"

// WriteLog TODO issue#173
type WriteLog struct {
	SchemaID     domain.ID
	InputID      domain.ID
	TemplateID   *domain.ID
	Identifier   domain.ID
	Code         uint16
	InputContext domain.InputContext
}
