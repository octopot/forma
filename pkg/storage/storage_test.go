package storage_test

import (
	"testing"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() { storage.Must(func(*storage.Storage) error { return nil }) })
	})
	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() { storage.Must(func(*storage.Storage) error { return errors.Simple("test") }) })
	})
}
