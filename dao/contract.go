package dao

import (
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// TODO v2: implement
// - add support MySQL

// TODO v3: implement
// MongoDB

// DataLayer defines behavior of database-agnostic data access object.
type DataLayer interface {
	// AddData inserts form data and returns its ID or an error if something went wrong.
	AddData(uuid data.UUID, values map[string][]string) (int64, error)
	// Schema would return a form schema with provided UUID or an error if something went wrong.
	Schema(uuid data.UUID) (form.Schema, error)
}

// Driver defines behavior of database-specific query executors.
type Driver interface {
	DataLayer
	// Dialect returns supported database dialect.
	Dialect() string
}

// DriverManager defines behavior of ...
type DriverManager interface {
}
