package postgres_test

import (
	"context"
	"testing"

	"github.com/kamilsk/form-api/pkg/storage/driver/postgres"
)

func TestNewUserContext(t *testing.T) {
	ctx := context.Background()
	_, _ = postgres.NewUserContext(nil, ctx)
	t.Run("token", func(t *testing.T) {
		// TODO
	})
}
