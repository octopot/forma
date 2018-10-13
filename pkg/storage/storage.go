package storage

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/executor"
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

// Database returns database configurator.
func Database(cnf config.DatabaseConfig) Configurator {
	return func(instance *Storage) (err error) {
		defer errors.Recover(&err)
		instance.exec = executor.New(cnf.DriverName())
		instance.db, err = sql.Open(cnf.DriverName(), string(cnf.DSN))
		if err == nil {
			instance.db.SetMaxOpenConns(cnf.MaxOpen)
			instance.db.SetMaxIdleConns(cnf.MaxIdle)
			instance.db.SetConnMaxLifetime(cnf.MaxLifetime)
		}
		return
	}
}

// Configurator defines a function which can use to configure the Storage.
type Configurator func(*Storage) error

// Storage is an implementation of Data Access Object.
type Storage struct {
	db   *sql.DB
	exec *executor.Executor
}

// Database returns the current database handle.
func (storage *Storage) Database() *sql.DB {
	return storage.db
}

// Dialect returns a supported database dialect.
func (storage *Storage) Dialect() string {
	return storage.exec.Dialect()
}

func (storage *Storage) connection(ctx context.Context) (*sql.Conn, func() error, error) {
	conn, err := storage.db.Conn(ctx)
	if err != nil {
		return conn, nil, errors.Database(errors.ServerErrorMessage, err, "trying to get connection")
	}
	return conn, conn.Close, nil
}
