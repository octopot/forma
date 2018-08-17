package query

import (
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
)

// InputFilter TODO
type InputFilter struct {
	SchemaID domain.ID
	From     time.Time
	To       time.Time
}

// WriteInput TODO
type WriteInput struct {
	SchemaID     domain.ID
	VerifiedData domain.InputData
}
