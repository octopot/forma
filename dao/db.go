package dao

import (
	"database/sql"
	"net/url"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
	"github.com/pkg/errors"
)

// Configurator defines a function which can use to configure DAO service.
type Configurator func(*service) error

// New returns a new instance of DAO service.
func New(configs ...Configurator) (*service, error) {
	instance := &service{}
	for _, configure := range configs {
		if err := configure(instance); err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Connection returns database connection Configurator.
func Connection(dsn *url.URL) Configurator {
	return func(instance *service) error {
		var err error
		instance.conn, err = sql.Open("postgres", dsn.String())
		return err
	}
}

type service struct {
	conn *sql.DB
}

// Schema returns form schema with provided UUID.
func (s *service) Schema(uuid data.UUID) (form.Schema, error) {
	var (
		schema form.Schema
		xml    []byte
	)
	row := s.conn.QueryRow(`SELECT schema FROM form_schema WHERE uuid = $1`, uuid)
	if err := row.Scan(&xml); err != nil {
		return schema, errors.WithMessage(err, "trying to find schema with UUID "+uuid.String())
	}
	if err := (&schema).UnmarshalFrom(xml); err != nil {
		return schema, errors.WithMessage(err, "trying to unmarshal schema with UUID "+uuid.String())
	}
	schema.ID = uuid.String()
	return schema, nil
}
