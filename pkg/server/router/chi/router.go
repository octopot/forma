package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	common "github.com/kamilsk/form-api/pkg/server/middleware"
	"github.com/kamilsk/form-api/pkg/server/router"
	internal "github.com/kamilsk/form-api/pkg/server/router/chi/middleware"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api router.Server) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	notImplemented := func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) }

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/{ID}", func(r chi.Router) {
			r.Use(internal.Pack("ID", "id"))
			r.Get("/", common.Encoder(api.GetV1))
			r.Post("/", api.Input)
		})
	})
	r.Route("/api/v2", func(r chi.Router) {
		r.Route("/schema", func(r chi.Router) {
			r.Route("/{ID}", func(r chi.Router) {
				r.Use(internal.Pack("ID", "id"))
				r.Get("/", notImplemented)
				r.Post("/", notImplemented)
			})
		})
		r.Route("/template", func(r chi.Router) {
			r.Route("/{ID}", func(r chi.Router) {
				r.Use(internal.Pack("ID", "id"))
				r.Get("/", notImplemented)
			})
		})
	})
	r.Route("/schema/{SCHEMA}/template/{TEMPLATE}", func(r chi.Router) {
		r.Use(internal.Pack("SCHEMA", "schema"))
		r.Use(internal.Pack("TEMPLATE", "template"))
		r.Get("/", notImplemented)
	})

	return r
}
