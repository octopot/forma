//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package dao -destination $PWD/dao/mock_database.go database/sql/driver Conn,Driver,Stmt,Rows
package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
		instance.conn, err = sql.Open(dsn.Scheme, dsn.String())
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
	row := s.conn.QueryRow(`SELECT "schema" FROM "form_schema" WHERE "uuid" = $1 AND "status" = 'enabled'`, uuid)
	if err := row.Scan(&xml); err != nil {
		return schema, errors.WithMessage(err, fmt.Sprintf("trying to find schema with UUID %q", uuid))
	}
	if err := (&schema).UnmarshalFrom(xml); err != nil {
		return schema, errors.WithMessage(err, fmt.Sprintf("trying to unmarshal schema with UUID %q from XML", uuid))
	}
	schema.ID = uuid.String()
	return schema, nil
}

// AddData inserts form data and returns its ID.
func (s *service) AddData(uuid data.UUID, values url.Values) (int64, error) {
	encoded, err := json.Marshal(values)
	if err != nil {
		return 0, errors.WithMessage(err, fmt.Sprintf("trying to marshal data into JSON with schema of %q", uuid))
	}
	result, err := s.conn.Exec(`INSERT INTO "form_data" ("uuid", "data") VALUES ($1, $2)`, uuid, encoded)
	if err != nil {
		return 0, errors.WithMessage(err, fmt.Sprintf("trying to insert JSON `%+v` with schema of %q", encoded, uuid))
	}
	return result.LastInsertId()
}
