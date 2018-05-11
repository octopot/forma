package middleware

import (
	"context"
	"net/http"

	"github.com/golang/gddo/httputil"
	"github.com/kamilsk/form-api/transfer/encoding"
)

// Notes:
// - related issue https://github.com/golang/go/issues/19307

// Encoder injects required response encoder to the request context.
func Encoder(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var contentType string
		if req.Header.Get("Accept") == "" {
			contentType = encoding.XML
		} else {
			contentType = httputil.NegotiateContentType(req, encoding.Offers(), "")
		}
		if !encoding.IsSupported(contentType) {
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}
		next(rw, req.WithContext(context.WithValue(req.Context(), EncoderKey{}, encoding.NewEncoder(rw, contentType))))
	}
}
