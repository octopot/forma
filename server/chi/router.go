package chi

import (
	"context"
	"net/http"
	"strings"

	"fmt"

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
				supported := map[string]string{
					"json": "application/json",
					//"toml": "application/toml",
					//"xml":  "application/xml",
					//"yaml": "application/yaml",
				}
				format := strings.ToLower(req.URL.Query().Get("format"))
				if format == "" {
					format = "json"
				}
				ct, ok := supported[format]
				if !ok {
					http.Error(rw, "Unsupported output format.", http.StatusUnsupportedMediaType)
					return
				}
				rw.Header().Set("Content-Type", ct)
				fmt.Fprintf(rw, "get form schema in format %q with Content-Type %q", format, ct)
			})
			r.Post("/", func(rw http.ResponseWriter, req *http.Request) {
				rw.Write([]byte("send form data"))
			})
		})
	})

	r.Route("/admin", func(r chi.Router) {})

	return r
}
