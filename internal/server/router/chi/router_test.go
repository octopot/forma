package chi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/ecosystem/forma/internal/server/router"
	"go.octolab.org/ecosystem/forma/internal/server/router/chi"
)

func TestNewRouter(t *testing.T) {
	type server struct{ Server }
	assert.NotPanics(t, func() { _ = chi.NewRouter(server{}) })
}
