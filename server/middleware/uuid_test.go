package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kamilsk/form-api/domen"
	"github.com/kamilsk/form-api/server/middleware"
	"github.com/stretchr/testify/assert"
)

const UUID domen.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestSchema(t *testing.T) {
	tests := []struct {
		name string
		uuid domen.UUID
		next func(uuid domen.UUID) (*domen.UUID, http.Handler)
		code int
	}{
		{"invalid uuid", "abc-def-ghi", func(uuid domen.UUID) (*domen.UUID, http.Handler) {
			return &uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			})
		}, http.StatusBadRequest},
		{"valid uuid", UUID, func(_ domen.UUID) (*domen.UUID, http.Handler) {
			uuid := new(domen.UUID)
			return uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				*uuid = req.Context().Value(middleware.SchemaKey{}).(domen.UUID)
			})
		}, http.StatusOK},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			rw, req := httptest.NewRecorder(), &http.Request{}
			uuid, next := tc.next(tc.uuid)
			middleware.Schema(tc.uuid.String(), rw, req, next)

			assert.Equal(t, tc.code, rw.Code)
			assert.Equal(t, tc.uuid, *uuid)
		})
	}
}

func TestTemplate(t *testing.T) {
	tests := []struct {
		name string
		uuid domen.UUID
		next func(uuid domen.UUID) (*domen.UUID, http.Handler)
		code int
	}{
		{"invalid uuid", "abc-def-ghi", func(uuid domen.UUID) (*domen.UUID, http.Handler) {
			return &uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			})
		}, http.StatusBadRequest},
		{"valid uuid", UUID, func(_ domen.UUID) (*domen.UUID, http.Handler) {
			uuid := new(domen.UUID)
			return uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				*uuid = req.Context().Value(middleware.TemplateKey{}).(domen.UUID)
			})
		}, http.StatusOK},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			rw, req := httptest.NewRecorder(), &http.Request{}
			uuid, next := tc.next(tc.uuid)
			middleware.Template(tc.uuid.String(), rw, req, next)

			assert.Equal(t, tc.code, rw.Code)
			assert.Equal(t, tc.uuid, *uuid)
		})
	}
}
