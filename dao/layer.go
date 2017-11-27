//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package dao -destination $PWD/dao/mock_db.go database/sql/driver Conn,Driver,Stmt,Rows
package dao

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
	"github.com/kamilsk/form-api/errors"
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

// Dialect returns supported database SQL dialect.
func (l *layer) Dialect() string {
	return "postgres"
}

// Schema would return a form schema with provided UUID or an error if something went wrong.
func (l *layer) Schema(uuid data.UUID) (form.Schema, error) {
	var (
		schema form.Schema
		raw    []byte
	)
	row := l.conn.QueryRow(`SELECT "schema" FROM "form_schema" WHERE "uuid" = $1 AND "status" = 'enabled'`, uuid)
	if err := row.Scan(&raw); err != nil {
		if err == sql.ErrNoRows {
			return schema, errors.NotFound(err, "schema with UUID %q not found", uuid)
		}
		return schema, errors.Database(err, "trying to populate schema with UUID %q", uuid)
	}
	if err := xml.Unmarshal(raw, &schema); err != nil {
		return schema, errors.Serialization(err, "trying to unmarshal schema with UUID %q from XML", uuid)
	}
	schema.ID = uuid.String()
	return schema, nil
}

// AddData inserts form data and returns its ID or an error if something went wrong.
func (l *layer) AddData(uuid data.UUID, values map[string][]string) (int64, error) {
	encoded, err := json.Marshal(values)
	if err != nil {
		return 0, errors.Serialization(err, "trying to marshal data into JSON with schema of %q", uuid)
	}
	result, err := l.conn.Exec(`INSERT INTO "form_data" ("uuid", "data") VALUES ($1, $2)`, uuid, encoded)
	if err != nil {
		return 0, errors.Database(err, "trying to insert JSON `%s` with schema of %q", encoded, uuid)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Database(err, "trying to get last insert ID of JSON `%s` with schema of %q", encoded, uuid)
	}
	return id, nil
}
