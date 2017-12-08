package dao

import (
	"database/sql"

	"github.com/kamilsk/form-api/dao/postgres"
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// Configurator defines a function which can use to configure DAO layer.
type Configurator func(*layer) error

// Must returns a new instance of DAO layer or panics if it cannot configure it.
func Must(configs ...Configurator) *layer {
	instance, err := New(configs...)
	if err != nil {
		panic(err)
	}
	return instance
}

// New returns a new instance of DAO layer or an error if it cannot configure it.
func New(configs ...Configurator) (*layer, error) {
	instance := &layer{}
	for _, configure := range configs {
		if err := configure(instance); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Connection returns database connection Configurator.
func Connection(driver, dsn string) Configurator {
	return func(instance *layer) error {
		var err error
		instance.conn, err = sql.Open(driver, dsn)
		return err
	}
}

type layer struct {
	conn *sql.DB
}

// Connection returns current database connection.
func (l *layer) Connection() *sql.DB {
	return l.conn
}

// AddData inserts form data and returns its ID or an error if something went wrong.
func (l *layer) AddData(uuid data.UUID, verified map[string][]string) (int64, error) {
	return postgres.AddData(l.conn, uuid, verified)
}

// Dialect returns supported database SQL dialect.
func (l *layer) Dialect() string {
	return postgres.Dialect()
}

// Schema would return a form schema with provided UUID or an error if something went wrong.
func (l *layer) Schema(uuid data.UUID) (form.Schema, error) {
	return postgres.Schema(l.conn, uuid)
}
