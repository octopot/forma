package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/encoder"
	"github.com/kamilsk/form-api/server/errors"
)

// Encoder injects required response encoder to the request context.
func Encoder(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		// Accept: text/html
		// Accept: image/*
		// Accept: text/html, application/xhtml+xml, application/xml; q=0.9, */*; q=0.8
		accept := defaultStringValue(req.Header.Get("Accept"), encoder.XML)
		contentType := strings.TrimSpace(strings.Split(strings.Split(accept, ";")[0], ",")[0])
		if !encoder.Support(contentType) {
			errors.NotSupportedContentType(encoder.Supported()).MarshalTo(rw) //nolint: errcheck
			return
		}
		next(rw, req.WithContext(context.WithValue(req.Context(), EncoderKey{}, encoder.New(rw, contentType))))
	}
}

// ValidateUUID validates form schema UUID and injects it to the request context.
func ValidateUUID(formUUID string, rw http.ResponseWriter, req *http.Request, next http.Handler) {
	uuid := data.UUID(formUUID)
	if uuid.IsEmpty() {
		errors.NotProvidedUUID().MarshalTo(rw) //nolint: errcheck
		return
	}
	if !uuid.IsValid() {
		errors.InvalidUUID().MarshalTo(rw) //nolint: errcheck
		return
	}
	next.ServeHTTP(rw, req.WithContext(context.WithValue(req.Context(), UUIDKey{}, uuid)))
}

func defaultStringValue(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
