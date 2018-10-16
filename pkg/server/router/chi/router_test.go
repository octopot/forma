//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package chi_test -destination $PWD/pkg/server/router/chi/mock_server_test.go github.com/kamilsk/form-api/pkg/server/router Server
package chi_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi/middleware"
	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/form-api/pkg/server/router/chi"
)

const UUID domain.ID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestChiRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	middleware.DefaultLogger = func(next http.Handler) http.Handler { return next }

	var (
		api    = NewMockServer(ctrl)
		router = NewRouter(api)
	)

	tests := []struct {
		name string
		data func() *http.Request
		code int
	}{
		{"GET /api/v1/{ID}", func() *http.Request {
			request := &http.Request{Method: http.MethodGet, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v1/" + UUID.String(),
			}}
			api.EXPECT().
				GetV1(gomock.Any(), gomock.Any()).
				Do(func(rw http.ResponseWriter, _ *http.Request) { rw.WriteHeader(http.StatusOK) })
			return request
		}, http.StatusOK},
		{"POST /api/v1/{ID}", func() *http.Request {
			request := &http.Request{Method: http.MethodPost, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v1/" + UUID.String(),
			}}
			api.EXPECT().
				Input(gomock.Any(), gomock.Any()).
				Do(func(rw http.ResponseWriter, _ *http.Request) { rw.WriteHeader(http.StatusFound) })
			return request
		}, http.StatusFound},

		{"GET /api/v2/schema/{ID}", func() *http.Request {
			return &http.Request{Method: http.MethodGet, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/schema/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},
		{"POST /api/v2/schema/{ID}", func() *http.Request {
			return &http.Request{Method: http.MethodPost, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/schema/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},

		{"GET /api/v2/template/{ID}", func() *http.Request {
			return &http.Request{Method: http.MethodGet, URL: &url.URL{
				Scheme: "http", Host: "dev", Path: "/api/v2/template/" + UUID.String(),
			}}
		}, http.StatusNotImplemented},

		{"GET /schema/{SCM_ID}/template/{TPL_ID}", func() *http.Request {
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
