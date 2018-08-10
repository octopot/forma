package postgres_test

import (
	"context"
	"testing"

	"github.com/kamilsk/form-api/pkg/storage/driver/postgres"
)

func TestNewInputContext(t *testing.T) {
	ctx := context.Background()
	_, _ = postgres.NewInputContext(nil, ctx)
	t.Run("read by ID", func(t *testing.T) {
		// TODO
	})
	t.Run("read by filter", func(t *testing.T) {
		// TODO
	})
}
