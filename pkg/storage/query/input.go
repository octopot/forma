package query

import (
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
)

// InputFilter TODO issue#173
type InputFilter struct {
	SchemaID domain.ID
	From     time.Time
	To       time.Time
}

// WriteInput TODO issue#173
type WriteInput struct {
	SchemaID     domain.ID
	VerifiedData domain.InputData
}

// WriteLog TODO issue#173
type WriteLog struct {
	SchemaID     domain.ID
	InputID      domain.ID
	TemplateID   *domain.ID
	Identifier   string
	Code         uint16
	InputContext domain.InputContext
}
