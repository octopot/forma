package postgres

import (
	"database/sql"
	"encoding/xml"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
	"github.com/kamilsk/form-api/errors"
)

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
