//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package server_test -destination $PWD/server/mock_contract_test.go github.com/kamilsk/form-api/server Service
package server_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/domain"
	"github.com/kamilsk/form-api/errors"
	"github.com/kamilsk/form-api/server"
	"github.com/kamilsk/form-api/server/middleware"
	"github.com/kamilsk/form-api/transfer/api/v1"
	"github.com/kamilsk/form-api/transfer/encoding"
	"github.com/stretchr/testify/assert"
)

const (
	HOST = "http://form-api.dev/"
	FAKE = domain.UUID("41ca5e09-3ce2-0094-b108-3ecc257c6fa4")
	ZERO = domain.UUID("00000000-0000-4000-8000-000000000000")
	UUID = domain.UUID("41ca5e09-3ce2-4094-b108-3ecc257c6fa4")
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		service = NewMockService(ctrl)
	)

	tests := []struct {
		name     string
		args     struct{ baseURL, tplPath string }
		panicked bool
	}{
		{"base URL is invalid", struct{ baseURL, tplPath string }{baseURL: "http://192.168.0.%31/"}, true},
		{"successful instance", struct{ baseURL, tplPath string }{baseURL: HOST, tplPath: "/"}, false},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if tc.panicked {
				assert.Panics(t, func() { server.New(tc.args.baseURL, tc.args.tplPath, service) })
			} else {
				assert.NotNil(t, server.New(tc.args.baseURL, tc.args.tplPath, service))
			}
		})
	}
}

func TestServer_GetV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		service = NewMockService(ctrl)
	)

	srv := server.New(HOST, "", service)

	tests := []struct {
		name    string
		request func(io.Writer) *http.Request
		code    int
	}{
		{http.StatusText(http.StatusInternalServerError), func(out io.Writer) *http.Request {
			req, _ := http.NewRequest(http.MethodGet, HOST, nil)
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.SchemaKey{}, FAKE))
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.EncoderKey{}, encoding.NewEncoder(out, encoding.XML)))
			service.EXPECT().
				HandleGetV1(v1.GetRequest{UUID: FAKE}).
				Return(v1.GetResponse{Error: errors.Database("", nil, "")})
			return req
		}, http.StatusInternalServerError},
		{http.StatusText(http.StatusNotFound), func(out io.Writer) *http.Request {
			req, _ := http.NewRequest(http.MethodGet, HOST, nil)
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.SchemaKey{}, ZERO))
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.EncoderKey{}, encoding.NewEncoder(out, encoding.XML)))
			service.EXPECT().
				HandleGetV1(v1.GetRequest{UUID: ZERO}).
				Return(v1.GetResponse{Error: errors.NotFound("", nil, "")})
			return req
		}, http.StatusNotFound},
		{http.StatusText(http.StatusOK), func(out io.Writer) *http.Request {
			req, _ := http.NewRequest(http.MethodGet, HOST, nil)
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.SchemaKey{}, UUID))
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.EncoderKey{}, encoding.NewEncoder(out, encoding.XML)))
			service.EXPECT().
				HandleGetV1(v1.GetRequest{UUID: UUID}).
				Return(v1.GetResponse{})
			return req
		}, http.StatusOK},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			srv.GetV1(rw, tc.request(rw))
			assert.Equal(t, tc.code, rw.Code)
		})
	}
}
