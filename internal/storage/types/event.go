package types

import (
	"time"

	"go.octolab.org/ecosystem/forma/internal/domain"
)

// Event TODO issue#173
type Event struct {
	ID         uint64              `db:"id"`
	AccountID  domain.ID           `db:"account_id"`
	SchemaID   domain.ID           `db:"schema_id"`
	InputID    domain.ID           `db:"input_id"`
	TemplateID *domain.ID          `db:"template_id"`
	Identifier *domain.ID          `db:"identifier"`
	Context    domain.InputContext `db:"context"`
	Code       int                 `db:"code"`
	URL        string              `db:"url"`
	CreatedAt  time.Time           `db:"created_at"`
	Account    *Account            `db:"-"`
	Schema     *Schema             `db:"-"`
	Input      *Input              `db:"-"`
	Template   *Template           `db:"-"`
}
