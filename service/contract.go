package service

import (
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// DataLayer defines behavior of DAO.
type DataLayer interface {
	Schema(data.UUID) (form.Schema, error)
	AddData(data.UUID, map[string][]string) (int64, error)
}
