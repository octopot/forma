package middleware

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Pack gets params from URL and puts it to a query.
func Pack(args ...string) func(http.Handler) http.Handler {
	if len(args) == 0 || len(args)%2 != 0 {
		panic("passed arguments must contain from => to list")
	}
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if err := req.ParseForm(); err == nil {
				for i, steps := 0, len(args)/2; i < steps; i++ {
					from, to := args[i], args[i+1]
					req.Form.Set(to, chi.URLParam(req, from))
				}
			}
			handler.ServeHTTP(rw, req)
		})
	}
}
