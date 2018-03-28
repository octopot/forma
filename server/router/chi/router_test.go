//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package chi_test -destination $PWD/server/router/chi/mock_contract_test.go github.com/kamilsk/form-api/server/router Server
package chi_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi/middleware"
	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/domain"
	"github.com/kamilsk/form-api/server/router/chi"
	"github.com/stretchr/testify/assert"
)

const UUID domain.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestChiRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	middleware.DefaultLogger = func(next http.Handler) http.Handler { return next }

	var (
		api    = NewMockServer(ctrl)
		router = chi.NewRouter(api)
	)

	tests := []struct {
		name string
		data func() *http.Request
		code int
	}{
		{"POST /api/v1", func() *http.Request {
			return &http.Request{Method: http.MethodPost, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v1",
			}}
		}, http.StatusNotImplemented},
		{"GET /api/v1/{UUID}", func() *http.Request {
			request := &http.Request{Method: http.MethodGet, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v1/" + UUID.String(),
			}}
			api.EXPECT().
				GetV1(gomock.Any(), gomock.Any()).
				Do(func(rw http.ResponseWriter, _ *http.Request) { rw.WriteHeader(http.StatusOK) })
			return request
		}, http.StatusOK},
		{"PUT /api/v1/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodPut, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v1/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
		{"DELETE /api/v1/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodDelete, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v1/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
		{"POST /api/v1/{UUID}", func() *http.Request {
			request := &http.Request{Method: http.MethodPost, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v1/" + UUID.String(),
			}}
			api.EXPECT().
				PostV1(gomock.Any(), gomock.Any()).
				Do(func(rw http.ResponseWriter, _ *http.Request) { rw.WriteHeader(http.StatusFound) })
			return request
		}, http.StatusFound},

		{"POST /api/v2/schema", func() *http.Request {
			return &http.Request{Method: http.MethodPost, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/schema",
			}}
		}, http.StatusNotImplemented},
		{"GET /api/v2/schema/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodGet, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/schema/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
		{"PUT /api/v2/schema/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodPut, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/schema/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
		{"DELETE /api/v2/schema/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodDelete, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/schema/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
		{"POST /api/v2/schema/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodPost, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/schema/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},

		{"POST /api/v2/template", func() *http.Request {
			return &http.Request{Method: http.MethodPost, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/template",
			}}
		}, http.StatusNotImplemented},
		{"GET /api/v2/template/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodGet, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/template/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
		{"PUT /api/v2/template/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodPut, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/template/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
		{"DELETE /api/v2/template/{UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodDelete, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/template/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},

		{"GET /schema/{SCM_UUID}/template/{TPL_UUID}", func() *http.Request {
			return &http.Request{Method: http.MethodGet, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/schema/" + UUID.String() + "/template/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, tc.data())

			assert.Equal(t, tc.code, rec.Code)
		})
	}
}
