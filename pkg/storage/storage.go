package storage

import (
	"context"
	"database/sql"
	"encoding/xml"
	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/executor"
	"github.com/kamilsk/form-api/pkg/storage/query"
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

// Database returns Database Configurator.
func Database(cnf config.DBConfig) Configurator {
	return func(instance *Storage) (err error) {
		defer errors.Recover(&err)
		instance.db, err = sql.Open(cnf.DriverName(), string(cnf.DSN))
		if err == nil {
			instance.db.SetMaxOpenConns(cnf.MaxOpen)
			instance.db.SetMaxIdleConns(cnf.MaxIdle)
			instance.db.SetConnMaxLifetime(cnf.MaxLifetime)
			instance.exec = executor.New(cnf.DriverName())
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

// TODO legacy

// AddData inserts a form data and returns its ID.
func (storage *Storage) AddData(ctx context.Context, schemaID domain.ID, verified domain.InputData) (domain.ID, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return "", err
	}
	defer closer()

	writer := storage.exec.InputWriter(ctx, conn)
	entity, err := writer.Write(query.WriteInput{SchemaID: string(schemaID), VerifiedData: verified})
	if err != nil {
		return "", err
	}
	return domain.ID(entity.ID), nil
}

// Schema returns the form schema by provided ID.
func (storage *Storage) Schema(ctx context.Context, id domain.ID) (domain.Schema, error) {
	var schema domain.Schema

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return schema, err
	}
	defer closer()

	reader := storage.exec.SchemaReader(ctx, conn)
	entity, err := reader.ReadByID(string(id))
	if err != nil {
		return schema, err
	}
	if decodeErr := xml.Unmarshal([]byte(entity.Definition), &schema); decodeErr != nil {
		return schema, errors.Serialization(errors.NeutralMessage, decodeErr,
			"trying to unmarshal the schema %q from XML `%s`", entity.ID, entity.Definition)
	}
	schema.Language, schema.Title = entity.Language, entity.Title
	return schema, nil
}

// Template returns the form template by provided ID.
func (storage *Storage) Template(ctx context.Context, id domain.ID) (domain.Template, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return "", err
	}
	defer closer()

	reader := storage.exec.TemplateReader(ctx, conn)
	entity, err := reader.ReadByID(string(id))
	if err != nil {
		return "", err
	}
	return domain.Template(entity.Definition), nil
}
