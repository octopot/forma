package postgres_test

import (
	"context"
	"testing"

	"github.com/kamilsk/form-api/pkg/storage/executor/internal/postgres"
)

func TestNewSchemaContext(t *testing.T) {
	ctx := context.Background()
	_ = postgres.NewSchemaContext(ctx, nil)
	t.Run("create", func(t *testing.T) {
		// TODO
	})
	t.Run("read", func(t *testing.T) {
		// TODO
	})
	t.Run("update", func(t *testing.T) {
		// TODO
	})
	t.Run("delete", func(t *testing.T) {
		// TODO
	})
}
