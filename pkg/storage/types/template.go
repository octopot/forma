package types

import (
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Template TODO issue#173
type Template struct {
	ID         domain.ID       `db:"id"`
	AccountID  domain.ID       `db:"account_id"`
	Title      string          `db:"title"`
	Definition domain.Template `db:"definition"`
	CreatedAt  time.Time       `db:"created_at"`
	UpdatedAt  *time.Time      `db:"updated_at"`
	DeletedAt  *time.Time      `db:"deleted_at"`
	Account    *Account        `db:"-"`
}
