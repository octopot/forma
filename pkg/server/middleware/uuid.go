package middleware

import (
	"context"
	"net/http"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Schema validates the passed Schema ID and injects it to the request context.
func Schema(uuid string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	if !domain.ID(uuid).IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next.ServeHTTP(rw, req.WithContext(context.WithValue(req.Context(), SchemaKey{}, domain.ID(uuid))))
}

// Template validates template ID and injects it to the request context.
func Template(uuid string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	if !domain.ID(uuid).IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	next.ServeHTTP(rw, req.WithContext(context.WithValue(req.Context(), TemplateKey{}, domain.ID(uuid))))
}
