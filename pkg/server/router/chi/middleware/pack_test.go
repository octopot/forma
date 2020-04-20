package middleware_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/form-api/pkg/server/router/chi/middleware"
)

const uuid = "10000000-2000-4000-8000-160000000004"

func TestPack(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		to       string
		ctx      func() *chi.Context
		expected string
	}{
		{"present", "ID", "id", func() *chi.Context {
			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("ID", uuid)
			return ctx
		}, uuid},
		{"not present", "ID", "id", chi.NewRouteContext, ""},
	}
	for _, test := range tests {
		tc := test
		t.Run("usual: "+test.name, func(t *testing.T) {
			ctx := context.WithValue(context.TODO(), chi.RouteCtxKey, tc.ctx())
			handler := Pack(tc.from, tc.to)(
				http.HandlerFunc(func(_ http.ResponseWriter, req *http.Request) {
					assert.Equal(t, tc.expected, req.Form.Get(tc.to))
				}),
			)
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/{%s}", tc.from), nil)
			assert.NoError(t, err)
			handler.ServeHTTP(nil, req.WithContext(ctx))
		})
	}

	panics := []struct {
		name string
		args []string
	}{
		{"nil args", nil},
		{"empty args", []string{}},
		{"odd args", []string{"ID", "id", "URL"}},
	}
	for _, test := range panics {
		tc := test
		t.Run("panic: "+test.name, func(t *testing.T) {
			assert.Panics(t, func() { Pack(tc.args...) })
		})
	}
}
