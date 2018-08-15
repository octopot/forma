package postgres_test

import (
	"context"
	"testing"

	"github.com/kamilsk/form-api/pkg/storage/executor/internal/postgres"
)

func TestNewUserContext(t *testing.T) {
	ctx := context.Background()
	_ = postgres.NewUserContext(ctx, nil)
	t.Run("token", func(t *testing.T) {
		// TODO
	})
}
