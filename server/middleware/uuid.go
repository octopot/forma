package middleware

import (
	"context"
	"net/http"

	"github.com/kamilsk/form-api/domen"
)

// Schema validates form schema UUID and injects it to the request context.
func Schema(uuid string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	if !domen.UUID(uuid).IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next.ServeHTTP(rw, req.WithContext(context.WithValue(req.Context(), SchemaKey{}, domen.UUID(uuid))))
}

// Template validates template UUID and injects it to the request context.
func Template(uuid string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	if !domen.UUID(uuid).IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next.ServeHTTP(rw, req.WithContext(context.WithValue(req.Context(), TemplateKey{}, domen.UUID(uuid))))
}
