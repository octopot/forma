package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/server/middleware"
	"github.com/stretchr/testify/assert"
)

const UUID domain.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestSchema(t *testing.T) {
	tests := []struct {
		name string
		uuid domain.UUID
		next func(uuid domain.UUID) (*domain.UUID, http.Handler)
		code int
	}{
		{"invalid uuid", "abc-def-ghi", func(uuid domain.UUID) (*domain.UUID, http.Handler) {
			return &uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			})
		}, http.StatusBadRequest},
		{"valid uuid", UUID, func(_ domain.UUID) (*domain.UUID, http.Handler) {
			uuid := new(domain.UUID)
			return uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				*uuid = req.Context().Value(middleware.SchemaKey{}).(domain.UUID)
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
		uuid domain.UUID
		next func(uuid domain.UUID) (*domain.UUID, http.Handler)
		code int
	}{
		{"invalid uuid", "abc-def-ghi", func(uuid domain.UUID) (*domain.UUID, http.Handler) {
			return &uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
			})
		}, http.StatusBadRequest},
		{"valid uuid", UUID, func(_ domain.UUID) (*domain.UUID, http.Handler) {
			uuid := new(domain.UUID)
			return uuid, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.WriteHeader(http.StatusOK)
				*uuid = req.Context().Value(middleware.TemplateKey{}).(domain.UUID)
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
