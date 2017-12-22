package dao

import (
	"database/sql"

	"github.com/kamilsk/form-api/dao/postgres"
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
)

// TODO v2: refactoring
// - replace sql.DB by DriverManager
// - move Connection to Driver/DriverManager layer
// - do not use postgres package directly, use DriverManager instead

// Configurator defines a function which can use to configure data access object.
type Configurator func(*Layer) error

// Must returns a new instance of data access object or panics if it cannot configure it.
func Must(configs ...Configurator) *Layer {
	instance, err := New(configs...)
	if err != nil {
		panic(err)
	}
	return instance
}

// New returns a new instance of data access object or an error if it cannot configure it.
func New(configs ...Configurator) (*Layer, error) {
	instance := &Layer{}
	for _, configure := range configs {
		if err := configure(instance); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Connection returns database connection Configurator.
func Connection(driver, dsn string) Configurator {
	return func(instance *Layer) error {
		var err error
		instance.conn, err = sql.Open(driver, dsn)
		return err
	}
}

// Layer is an implementation of Data Access Object.
type Layer struct {
	conn *sql.DB
}

// Connection returns current database connection.
func (l *Layer) Connection() *sql.DB {
	return l.conn
}

// AddData inserts form data and returns its ID or an error if something went wrong.
func (l *Layer) AddData(uuid data.UUID, verified map[string][]string) (int64, error) {
	return postgres.AddData(l.conn, uuid, verified)
}

// Dialect returns supported database SQL dialect.
func (l *Layer) Dialect() string {
	return postgres.Dialect()
}

// Schema would return a form schema with provided UUID or an error if something went wrong.
func (l *Layer) Schema(uuid data.UUID) (form.Schema, error) {
	return postgres.Schema(l.conn, uuid)
}
