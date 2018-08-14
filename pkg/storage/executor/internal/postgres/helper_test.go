package postgres_test

import (
	"time"

	"github.com/kamilsk/form-api/pkg/storage/query"
)

func Token() *query.Token {
	return &query.Token{
		ID:        "10000000-2000-4000-8000-160000000003",
		UserID:    "10000000-2000-4000-8000-160000000002",
		ExpiredAt: nil,
		CreatedAt: time.Now(),
		User: &query.User{
			ID:        "10000000-2000-4000-8000-160000000002",
			AccountID: "10000000-2000-4000-8000-160000000001",
			Name:      "Demo user",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
			DeletedAt: nil,
			Account: &query.Account{
				ID:        "10000000-2000-4000-8000-160000000001",
				Name:      "Demo account",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
				DeletedAt: nil,
			},
		},
	}
}
