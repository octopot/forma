package query

import "time"

// InputFilter TODO
type InputFilter struct {
	SchemaID string
	From     time.Time
	To       time.Time
}
