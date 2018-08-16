package draft

import (
	"database/sql"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
)

type Signer interface {
	Token() string
}

type Encoder interface {
	Encode(string) string
	Decode(string) string
}

// UUID returns a new generated unique identifier.
func UUID(db *sql.DB) (domain.ID, error) {
	var id domain.ID
	row := db.QueryRow(`SELECT uuid_generate_v4()`)
	if err := row.Scan(&id); err != nil {
		return id, errors.Database(errors.ServerErrorMessage, err, "trying to populate UUID")
	}
	return id, nil
}
