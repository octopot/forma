package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/form-api/server"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api server.FormAPI) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/{UUID}", func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					server.ValidateUUID(chi.URLParam(req, "UUID"), rw, req, next)
				})
			})

			r.Get("/", api.GetV1)
			r.Post("/", api.PostV1)
		})
	})

	r.Route("/admin", func(r chi.Router) {})

	return r
}