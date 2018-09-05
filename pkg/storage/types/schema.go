package types

import (
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Schema TODO issue#173
type Schema struct {
	ID         domain.ID     `db:"id"`
	AccountID  domain.ID     `db:"account_id"`
	Title      string        `db:"title"`
	Definition domain.Schema `db:"definition"`
	CreatedAt  time.Time     `db:"created_at"`
	UpdatedAt  *time.Time    `db:"updated_at"`
	DeletedAt  *time.Time    `db:"deleted_at"`
	Account    *Account      `db:"-"`
}
