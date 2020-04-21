package types

import (
	"time"

	"go.octolab.org/ecosystem/forma/internal/domain"
)

// Input TODO issue#173
type Input struct {
	ID        domain.ID        `db:"id"`
	SchemaID  domain.ID        `db:"schema_id"`
	Data      domain.InputData `db:"data"`
	CreatedAt time.Time        `db:"created_at"`
	Schema    *Schema          `db:"-"`
}
