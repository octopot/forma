package postgres

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
)

const dialect = "postgres"

// Dialect returns supported database dialect.
func Dialect() string {
	return dialect
}

// AddData inserts form data and returns their ID.
func AddData(db *sql.DB, uuid domain.UUID, verified map[string][]string) (int64, error) {
	var id int64
	encoded, err := json.Marshal(verified)
	if err != nil {
		return id, errors.Serialization(errors.ServerErrorMessage, err,
			"trying to marshal data `%#v` into JSON with the schema %q", verified, uuid)
	}
	err = db.QueryRow(`INSERT INTO "form_data" ("uuid", "data") VALUES ($1, $2) RETURNING "id"`, uuid, encoded).
		Scan(&id)
	if err != nil {
		return id, errors.Database(errors.ServerErrorMessage, err,
			"trying to insert JSON `%s` with the schema %q", encoded, uuid)
	}
	return id, nil
}

// Schema returns the form schema with provided UUID.
func Schema(db *sql.DB, uuid domain.UUID) (domain.Schema, error) {
	var (
		schema domain.Schema
		blob   = [1024]byte{}
		raw    = blob[:0]
	)
	row := db.QueryRow(`SELECT "schema" FROM "form_schema" WHERE "uuid" = $1 AND "status" = 'enabled'`, uuid)
	if err := row.Scan(&raw); err != nil {
		if err == sql.ErrNoRows {
			return schema, errors.NotFound(errors.SchemaNotFoundMessage, err, "schema %q not found or disabled", uuid)
		}
		return schema, errors.Database(errors.ServerErrorMessage, err, "trying to populate schema %q", uuid)
	}
	if err := xml.Unmarshal(raw, &schema); err != nil {
		return schema, errors.Serialization(errors.NeutralMessage, err,
			"trying to unmarshal schema %q from XML `%s`", uuid, raw)
	}
	return schema, nil
}

// UUID returns a new generated unique identifier.
func UUID(db *sql.DB) (domain.UUID, error) {
	var id domain.UUID
	row := db.QueryRow(`SELECT uuid_generate_v4()`)
	if err := row.Scan(&id); err != nil {
		return id, errors.Database(errors.ServerErrorMessage, err, "trying to populate UUID")
	}
	return id, nil
}
