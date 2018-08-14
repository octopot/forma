package storage

import (
	"database/sql"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/storage/postgres"
	"github.com/pkg/errors"
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
		if err := errors.WithStack(configure(instance)); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Connection returns database connection Configurator.
func Connection(cnf config.DBConfig) Configurator {
	return func(instance *Storage) error {
		var err error
		instance.conn, err = sql.Open(cnf.DriverName(), string(cnf.DSN))
		if err == nil {
			instance.conn.SetMaxOpenConns(cnf.MaxOpen)
			instance.conn.SetMaxIdleConns(cnf.MaxIdle)
			instance.conn.SetConnMaxLifetime(cnf.MaxLifetime)
		}
		return err
	}
}

// Configurator defines a function which can use to configure the Storage.
type Configurator func(*Storage) error

// Storage is an implementation of Data Access Object.
type Storage struct {
	conn *sql.DB
}

//
// TODO refactoring
//

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

// UUID returns a new generated unique identifier.
func (l *Storage) UUID() (domain.UUID, error) {
	return postgres.UUID(l.conn)
}
