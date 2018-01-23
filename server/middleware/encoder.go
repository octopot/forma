package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kamilsk/form-api/transfer/encoding"
)

// Encoder injects required response encoder to the request context.
func Encoder(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		// Accept: text/html
		// Accept: image/*
		// Accept: text/html, application/xhtml+xml, application/xml; q=0.9, */*; q=0.8
		accept := fallback(req.Header.Get("Accept"), encoding.XML)
		contentType := strings.TrimSpace(strings.Split(strings.Split(accept, ";")[0], ",")[0])
		if !encoding.IsSupported(contentType) {
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}
		next(rw, req.WithContext(context.WithValue(req.Context(), EncoderKey{}, encoding.NewEncoder(rw, contentType))))
	}
}

func fallback(value string, defaultValues ...string) string {
	if value == "" {
		for _, value := range defaultValues {
			if value != "" {
				return value
			}
		}
	}
	return value
}
