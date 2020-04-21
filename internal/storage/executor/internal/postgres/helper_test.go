package postgres_test

import (
	"time"

	"go.octolab.org/ecosystem/forma/internal/storage/types"
)

func DemoToken() *types.Token {
	var (
		account = types.Account{
			Name:      "Demo Account",
			CreatedAt: time.Now(),
		}
		user = types.User{
			AccountID: "10000000-2000-4000-8000-160000000001",
			Name:      "Demo User",
			CreatedAt: time.Now(),
		}
		token = types.Token{
			ID:        "10000000-2000-4000-8000-160000000003",
			UserID:    "10000000-2000-4000-8000-160000000002",
			CreatedAt: time.Now(),
		}
	)
	account.ID, user.ID = user.AccountID, token.UserID
	user.Account, token.User = &account, &user
	user.Tokens, account.Users = append(user.Tokens, &token), append(account.Users, &user)
	return &token
}
