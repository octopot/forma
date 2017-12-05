package postgres

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
	"github.com/kamilsk/form-api/errors"
)

// AddData inserts form data and returns its ID or an error if something went wrong.
func AddData(db *sql.DB, uuid data.UUID, values map[string][]string) (int64, error) {
	encoded, err := json.Marshal(values)
	if err != nil {
		return 0, errors.Serialization(err, "trying to marshal data into JSON with schema of %q", uuid)
	}
	result, err := db.Exec(`INSERT INTO "form_data" ("uuid", "data") VALUES ($1, $2)`, uuid, encoded)
	if err != nil {
		return 0, errors.Database(err, "trying to insert JSON `%s` with schema of %q", encoded, uuid)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Database(err, "trying to get last insert ID of JSON `%s` with schema of %q", encoded, uuid)
	}
	return id, nil
}

// Dialect returns supported database SQL dialect.
func Dialect() string {
	return "postgres"
}

// Schema would return a form schema with provided UUID or an error if something went wrong.
func Schema(db *sql.DB, uuid data.UUID) (form.Schema, error) {
	var (
		schema form.Schema
		raw    []byte
	)
	row := db.QueryRow(`SELECT "schema" FROM "form_schema" WHERE "uuid" = $1 AND "status" = 'enabled'`, uuid)
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
