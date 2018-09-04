package postgres_test

import (
	"time"

	"github.com/kamilsk/form-api/pkg/storage/types"
)

func Token() *types.Token {
	return &types.Token{
		ID:        "10000000-2000-4000-8000-160000000003",
		UserID:    "10000000-2000-4000-8000-160000000002",
		ExpiredAt: nil,
		CreatedAt: time.Now(),
		User: &types.User{
			ID:        "10000000-2000-4000-8000-160000000002",
			AccountID: "10000000-2000-4000-8000-160000000001",
			Name:      "Demo user",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
			DeletedAt: nil,
			Account: &types.Account{
				ID:        "10000000-2000-4000-8000-160000000001",
				Name:      "Demo account",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
				DeletedAt: nil,
			},
		},
	}
}
