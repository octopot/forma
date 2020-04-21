package postgres_test

import (
	"context"
	"testing"

	"go.octolab.org/ecosystem/forma/internal/storage/executor"
	. "go.octolab.org/ecosystem/forma/internal/storage/executor/internal/postgres"
)

func TestTemplateEditor(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TemplateEditor = NewTemplateContext(ctx, nil)
	})
	t.Run("read", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TemplateEditor = NewTemplateContext(ctx, nil)
	})
	t.Run("update", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TemplateEditor = NewTemplateContext(ctx, nil)
	})
	t.Run("delete", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TemplateEditor = NewTemplateContext(ctx, nil)
	})
}

func TestTemplateReader(t *testing.T) {
	t.Run("read by ID", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var _ executor.TemplateReader = NewTemplateContext(ctx, nil)
	})
}
