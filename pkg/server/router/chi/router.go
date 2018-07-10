package chi

import (
	"net/http"

	common "github.com/kamilsk/form-api/pkg/server/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/form-api/pkg/server/router"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api router.Server) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	notImplemented := func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) }

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/", notImplemented)

		r.Route("/{UUID}", func(r chi.Router) {
			r.Use(ctxPacker(common.Schema, "UUID"))

			r.Get("/", common.Encoder(api.GetV1))
			r.Put("/", notImplemented)
			r.Delete("/", notImplemented)

			r.Post("/", api.PostV1)
		})
	})

	r.Route("/api/v2", func(r chi.Router) {
		r.Route("/schema", func(r chi.Router) {
			r.Post("/", notImplemented)

			r.Route("/{UUID}", func(r chi.Router) {
				r.Use(ctxPacker(common.Schema, "UUID"))

				r.Get("/", notImplemented)
				r.Put("/", notImplemented)
				r.Delete("/", notImplemented)

				r.Post("/", notImplemented)
			})
		})

		r.Route("/template", func(r chi.Router) {
			r.Post("/", notImplemented)

			r.Route("/{UUID}", func(r chi.Router) {
				r.Use(ctxPacker(common.Template, "UUID"))

				r.Get("/", notImplemented)
				r.Put("/", notImplemented)
				r.Delete("/", notImplemented)
			})
		})
	})

	r.Route("/schema/{SCM_UUID}/template/{TPL_UUID}", func(r chi.Router) {
		r.Use(ctxPacker(common.Schema, "SCM_UUID"))
		r.Use(ctxPacker(common.Template, "TPL_UUID"))

		r.Get("/", notImplemented)
	})

	return r
}
