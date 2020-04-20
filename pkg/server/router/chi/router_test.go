package chi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/form-api/pkg/server/router"
	"github.com/kamilsk/form-api/pkg/server/router/chi"
)

func TestNewRouter(t *testing.T) {
	type server struct{ Server }
	assert.NotPanics(t, func() { _ = chi.NewRouter(server{}) })
}
