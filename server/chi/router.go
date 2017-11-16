package chi

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/form-api/server"
)

// NewRouter returns configured router.
func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/{UUID}", func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					next.ServeHTTP(rw,
						req.WithContext(
							context.WithValue(req.Context(), server.UUIDKey{}, chi.URLParam(req, "UUID"))))
				})
			})
			r.Use(server.UUID)

			r.Get("/", func(rw http.ResponseWriter, req *http.Request) {
				rw.Write([]byte("get form schema"))
			})
			r.Post("/", func(rw http.ResponseWriter, req *http.Request) {
				rw.Write([]byte("send form data"))
			})
		})
	})

	r.Route("/admin", func(r chi.Router) {})

	return r
}
