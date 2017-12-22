package chi

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/form-api/server"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api server.FormAPI, withProfiler bool) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	// TODO v2: support CRUD
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })

		r.Route("/{UUID}", func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					server.ValidateUUID(chi.URLParam(req, "UUID"), rw, req, next)
				})
			})

			r.Get("/", server.Encoder(api.GetV1))
			r.Post("/", api.PostV1)
			r.Put("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
			r.Delete("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
		})
	})

	// TODO v4: extended API
	r.Route("/api/v2", func(r chi.Router) {
		r.Route("/schema", func(r chi.Router) {
			r.Post("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })

			r.Route("/{UUID}", func(r chi.Router) {
				r.Get("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
				r.Post("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
				r.Put("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
				r.Delete("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
			})
		})

		r.Route("/template", func(r chi.Router) {
			r.Post("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })

			r.Route("/{UUID}", func(r chi.Router) {
				r.Get("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
				r.Post("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
				r.Put("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
				r.Delete("/", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) })
			})
		})
	})

	if withProfiler {
		r.Route("/debug/pprof", func(r chi.Router) {
			r.Get("/", pprof.Index)
			r.Get("/cmdline", pprof.Cmdline)
			r.Get("/profile", pprof.Profile)
			r.Get("/symbol", pprof.Symbol)
			r.Get("/trace", pprof.Trace)
		})
	}

	return r
}
