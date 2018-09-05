package storage

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/executor"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/storage/types"
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

// Schema returns the form schema by provided ID.
func (storage *Storage) Schema(ctx context.Context, id domain.ID) (domain.Schema, error) {
	var schema domain.Schema

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return schema, err
	}
	defer closer()

	entity, err := storage.exec.SchemaReader(ctx, conn).ReadByID(id)
	if err != nil {
		return schema, err
	}
	return entity.Definition, nil
}

// Template returns the form template by provided ID.
func (storage *Storage) Template(ctx context.Context, id domain.ID) (domain.Template, error) {
	var template domain.Template

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return template, err
	}
	defer closer()

	entity, err := storage.exec.TemplateReader(ctx, conn).ReadByID(id)
	if err != nil {
		return template, err
	}
	return entity.Definition, nil
}

// HandleInput TODO
func (storage *Storage) HandleInput(ctx context.Context, schemaID domain.ID, verified domain.InputData) (*types.Input, error) {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	entity, err := storage.exec.InputWriter(ctx, conn).Write(query.WriteInput{SchemaID: schemaID, VerifiedData: verified})
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// LogRequest TODO
func (storage *Storage) LogRequest(ctx context.Context, input *types.Input, meta domain.InputContext) error {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return err
	}
	defer closer()

	storage.exec.LogWriter(ctx, conn).Write(query.WriteLog{
		SchemaID:   input.SchemaID,
		InputID:    input.ID,
		TemplateID: input.Data.Template(),

		// TODO issue#171
		Identifier:   "",
		Code:         201,
		InputContext: meta,
	})
	return nil
}
