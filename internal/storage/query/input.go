package query

import (
	"time"

	"go.octolab.org/ecosystem/forma/internal/domain"
)

// InputFilter TODO issue#173
type InputFilter struct {
	SchemaID domain.ID
	From     *time.Time
	To       *time.Time
}

// WriteInput TODO issue#173
type WriteInput struct {
	SchemaID     domain.ID
	VerifiedData domain.InputData
}
