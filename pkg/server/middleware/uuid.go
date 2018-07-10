package middleware

import (
	"context"
	"net/http"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Schema validates the passed Schema ID and injects it to the request context.
func Schema(uuid string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	if !domain.UUID(uuid).IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next.ServeHTTP(rw, req.WithContext(context.WithValue(req.Context(), SchemaKey{}, domain.UUID(uuid))))
}

// Template validates template UUID and injects it to the request context.
func Template(uuid string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	if !domain.UUID(uuid).IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next.ServeHTTP(rw, req.WithContext(context.WithValue(req.Context(), TemplateKey{}, domain.UUID(uuid))))
}
