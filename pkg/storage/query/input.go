package query

import "time"

// InputFilter TODO
type InputFilter struct {
	SchemaID string
	From     time.Time
	To       time.Time
}

// WriteInput TODO
type WriteInput struct {
	SchemaID     string
	VerifiedData map[string][]string
}
