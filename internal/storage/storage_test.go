package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.octolab.org/ecosystem/forma/internal/errors"
	. "go.octolab.org/ecosystem/forma/internal/storage"
)

func TestMust(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert.NotPanics(t, func() { Must(func(*Storage) error { return nil }) })
	})
	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() { Must(func(*Storage) error { return errors.Simple("test") }) })
	})
}
