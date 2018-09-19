package types

import (
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Log TODO issue#173
type Log struct {
	ID         uint64              `db:"id"`
	AccountID  domain.ID           `db:"account_id"`
	SchemaID   domain.ID           `db:"schema_id"`
	InputID    domain.ID           `db:"input_id"`
	TemplateID *domain.ID          `db:"template_id"`
	Identifier domain.ID           `db:"identifier"`
	Code       uint16              `db:"code"`
	Context    domain.InputContext `db:"context"`
	CreatedAt  time.Time           `db:"created_at"`
	Account    *Account            `db:"-"`
	Schema     *Schema             `db:"-"`
	Input      *Input              `db:"-"`
	Template   *Template           `db:"-"`
}
