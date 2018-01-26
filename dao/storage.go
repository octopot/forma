package dao

import (
	"database/sql"

	"github.com/kamilsk/form-api/dao/postgres"
	"github.com/kamilsk/form-api/domain"
)

// Must returns a new instance of the Storage or panics if it cannot configure it.
func Must(configs ...Configurator) *Storage {
	instance, err := New(configs...)
	if err != nil {
		panic(err)
	}
	return instance
}

// New returns a new instance of the Storage or an error if it cannot configure it.
func New(configs ...Configurator) (*Storage, error) {
	instance := &Storage{}
	for _, configure := range configs {
		if err := configure(instance); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Connection returns database connection Configurator.
func Connection(driver, dsn string) Configurator {
	return func(instance *Storage) error {
		var err error
		instance.conn, err = sql.Open(driver, dsn)
		return err
	}
}

// Configurator defines a function which can use to configure the Storage.
type Configurator func(*Storage) error

// Storage is an implementation of Data Access Object.
type Storage struct {
	conn *sql.DB
}

// Connection returns current database connection.
func (l *Storage) Connection() *sql.DB {
	return l.conn
}

// Dialect returns supported database dialect.
func (l *Storage) Dialect() string {
	return postgres.Dialect()
}

// AddData inserts form data and returns their ID.
func (l *Storage) AddData(uuid domain.UUID, verified map[string][]string) (int64, error) {
	return postgres.AddData(l.conn, uuid, verified)
}

// Schema returns the form schema with provided UUID.
func (l *Storage) Schema(uuid domain.UUID) (domain.Schema, error) {
	return postgres.Schema(l.conn, uuid)
}
