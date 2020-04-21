package types

import (
	"time"

	"go.octolab.org/ecosystem/forma/internal/domain"
)

// Token TODO issue#173
type Token struct {
	ID        domain.ID  `db:"id"`
	UserID    domain.ID  `db:"user_id"`
	ExpiredAt *time.Time `db:"expired_at"`
	CreatedAt time.Time  `db:"created_at"`
	User      *User      `db:"-"`
}
