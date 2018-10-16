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
		r.Route("/{ID}", func(r chi.Router) {
			r.Use(ctxPacker(common.Schema, "ID"))
			r.Get("/", common.Encoder(api.GetV1))
			r.Post("/", api.HandleInput)
		})
	})
	r.Route("/api/v2", func(r chi.Router) {
		r.Route("/schema", func(r chi.Router) {
			r.Route("/{ID}", func(r chi.Router) {
				r.Use(ctxPacker(common.Schema, "ID"))
				r.Get("/", notImplemented)
				r.Post("/", notImplemented)
			})
		})
		r.Route("/template", func(r chi.Router) {
			r.Route("/{ID}", func(r chi.Router) {
				r.Use(ctxPacker(common.Template, "ID"))
				r.Get("/", notImplemented)
			})
		})
	})
	r.Route("/schema/{SCM_ID}/template/{TPL_ID}", func(r chi.Router) {
		r.Use(ctxPacker(common.Schema, "SCM_ID"))
		r.Use(ctxPacker(common.Template, "TPL_ID"))
		r.Get("/", notImplemented)
	})

	return r
}
