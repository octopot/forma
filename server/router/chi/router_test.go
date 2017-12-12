//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package chi_test -destination $PWD/server/router/chi/mock_contract_test.go github.com/kamilsk/form-api/server FormAPI
//go:generate mockgen -package chi_test -destination $PWD/server/router/chi/mock_http_test.go net/http ResponseWriter
package chi_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/server/router/chi"
)

const UUID data.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestChiRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	api := NewMockFormAPI(ctrl)
	api.EXPECT().GetV1(gomock.Any(), gomock.Any())
	rw := NewMockResponseWriter(ctrl)
	rw.EXPECT().Header().AnyTimes()
	r := chi.NewRouter(api, true)

	r.ServeHTTP(rw, &http.Request{Method: "GET", URL: &url.URL{
		Scheme: "http://",
		Host:   "dev",
		Path:   "/api/v1/" + UUID.String(),
	}})
}
