package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/form-api/server"
)

func main() {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
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
	log.Fatal(http.ListenAndServe(addr(), r))
}

func addr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return os.Getenv("BIND") + ":" + port
}
