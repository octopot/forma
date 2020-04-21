package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/ecosystem/forma/internal/server/middleware"
)

func TestPixel(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		next    http.HandlerFunc
		isPixel bool
	}{
		{"get pixel", "some.url?pixel", func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/plain")
			rw.WriteHeader(http.StatusOK)
		}, true},
		{"get text", "some.url", func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/plain")
			rw.WriteHeader(http.StatusOK)
		}, false},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			rw, req := httptest.NewRecorder(), func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, tc.query, nil)
				return req
			}()
			Pixel(tc.next)(rw, req)
			assert.Equal(t, tc.isPixel, "image/gif" == rw.Header().Get("Content-Type"))
		})
	}
}
