package types

import (
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Input TODO
type Input struct {
	ID        domain.ID        `db:"id"`
	SchemaID  domain.ID        `db:"schema_id"`
	Data      domain.InputData `db:"data"`
	CreatedAt time.Time        `db:"created_at"`
	Schema    *Schema          `db:"-"`
}
