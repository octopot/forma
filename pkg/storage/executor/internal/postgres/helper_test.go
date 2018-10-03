package postgres_test

import (
	"time"

	"github.com/kamilsk/form-api/pkg/storage/types"
)

func DemoToken() *types.Token {
	var (
		token = types.Token{
			ID:        "10000000-2000-4000-8000-160000000003",
			UserID:    "10000000-2000-4000-8000-160000000002",
			CreatedAt: time.Now(),
		}
		user = types.User{
			AccountID: "10000000-2000-4000-8000-160000000001",
			Name:      "Demo user",
			CreatedAt: time.Now(),
		}
		account = types.Account{
			Name:      "Demo account",
			CreatedAt: time.Now(),
		}
	)
	user.ID, account.ID = token.UserID, user.AccountID
	token.User, user.Account = &user, &account
	user.Tokens, account.Users = append(user.Tokens, &token), append(account.Users, &user)
	return &token
}
