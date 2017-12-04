package postgres

import (
	"database/sql"
	"encoding/json"

	"github.com/kamilsk/form-api/data"
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
