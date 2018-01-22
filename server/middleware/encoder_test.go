package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kamilsk/form-api/server/middleware"
	"github.com/kamilsk/form-api/transfer/encoding"
	"github.com/stretchr/testify/assert"
)

func TestEncoder(t *testing.T) {
	tests := []struct {
		name   string
		accept string
		next   func() (*encoding.Encoder, http.HandlerFunc)
		code   int
	}{
		{"empty header", "", func() (*encoding.Encoder, http.HandlerFunc) {
			encoder := new(encoding.Encoder)
			return encoder, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				*encoder = req.Context().Value(middleware.EncoderKey{}).(encoding.Encoder)
			})
		}, http.StatusOK},
		{"supported", "text/html", func() (*encoding.Encoder, http.HandlerFunc) {
			encoder := new(encoding.Encoder)
			return encoder, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				*encoder = req.Context().Value(middleware.EncoderKey{}).(encoding.Encoder)
			})
		}, http.StatusOK},
		{"not supported", "image/*", func() (*encoding.Encoder, http.HandlerFunc) {
			encoder := new(encoding.Encoder)
			*encoder = nil
			return encoder, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			})
		}, http.StatusNotAcceptable},
		{"complex", "text/html, application/xml; q=0.9, */*; q=0.8", func() (*encoding.Encoder, http.HandlerFunc) {
			encoder := new(encoding.Encoder)
			return encoder, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				*encoder = req.Context().Value(middleware.EncoderKey{}).(encoding.Encoder)
			})
		}, http.StatusOK},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			rw, req := httptest.NewRecorder(), &http.Request{Header: map[string][]string{"Accept": {tc.accept}}}
			encoder, next := tc.next()
			middleware.Encoder(next)(rw, req)

			assert.Equal(t, tc.code, rw.Code)
			assert.NotNil(t, encoder)
		})
	}
}
