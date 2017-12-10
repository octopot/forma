package service

import (
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// DataLayer defines behavior of DAO.
type DataLayer interface {
	// AddData inserts form data and returns its ID or an error if something went wrong.
	AddData(uuid data.UUID, values map[string][]string) (int64, error)
	// Schema would return a form schema with provided UUID or an error if something went wrong.
	Schema(uuid data.UUID) (form.Schema, error)
}
