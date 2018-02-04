package chi

import (
	"net/http"
	"net/http/pprof"

	common "github.com/kamilsk/form-api/server/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/form-api/server/router"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api router.Server, withProfiler bool) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	notImplemented := func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) }

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/", notImplemented)

		r.Route("/{UUID}", func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					common.Schema(chi.URLParam(req, "UUID"), rw, req, next)
				})
			})

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
				r.Use(func(next http.Handler) http.Handler {
					return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
						common.Schema(chi.URLParam(req, "UUID"), rw, req, next)
					})
				})

				r.Get("/", notImplemented)
				r.Put("/", notImplemented)
				r.Delete("/", notImplemented)

				r.Post("/", notImplemented)
			})
		})

		r.Route("/template", func(r chi.Router) {
			r.Post("/", notImplemented)

			r.Route("/{UUID}", func(r chi.Router) {
				r.Use(func(next http.Handler) http.Handler {
					return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
						common.Template(chi.URLParam(req, "UUID"), rw, req, next)
					})
				})

				r.Get("/", notImplemented)
				r.Put("/", notImplemented)
				r.Delete("/", notImplemented)
			})
		})
	})

	r.Route("/schema/{SCM_UUID}/template/{TPL_UUID}", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				common.Schema(chi.URLParam(req, "SCM_UUID"), rw, req, next)
			})
		})
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				common.Schema(chi.URLParam(req, "TPL_UUID"), rw, req, next)
			})
		})

		r.Get("/", notImplemented)
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
