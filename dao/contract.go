package dao

import (
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// Executor defines behavior of database specific SQL executor.
type Executor interface {
	// AddData inserts form data and returns its ID or an error if something went wrong.
	AddData(uuid data.UUID, values map[string][]string) (int64, error)
	// Dialect returns supported database SQL dialect.
	Dialect() string
	// Schema would return a form schema with provided UUID or an error if something went wrong.
	Schema(uuid data.UUID) (form.Schema, error)
}
