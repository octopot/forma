//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package server_test -destination $PWD/pkg/server/mock_contract_test.go github.com/kamilsk/form-api/pkg/server Service
package server_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/server"
	"github.com/kamilsk/form-api/pkg/server/middleware"
	"github.com/kamilsk/form-api/pkg/transfer/api/v1"
	"github.com/kamilsk/form-api/pkg/transfer/encoding"
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
				Return(v1.GetResponse{Schema: domain.Schema{Title: "schema", Inputs: []domain.Input{{Title: "input"}}}})
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

func TestServer_PostV1(t *testing.T) {
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
		{http.StatusText(http.StatusBadRequest) + ", missing form body", func(out io.Writer) *http.Request {
			req, _ := http.NewRequest(http.MethodPost, HOST, nil)
			return req
		}, http.StatusBadRequest},
		{http.StatusText(http.StatusBadRequest) + ", invalid input", func(out io.Writer) *http.Request {
			req, _ := http.NewRequest(http.MethodPost, HOST, strings.NewReader("email=invalid"))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.SchemaKey{}, UUID))
			data, err := func() (map[string][]string, error) {
				data := map[string][]string{"email": {"invalid"}}
				_, err := domain.Schema{Inputs: []domain.Input{{Name: "email", Type: domain.EmailType}}}.Validate(data)
				return data, errors.Validation("", err, "")
			}()
			service.EXPECT().
				HandlePostV1(v1.PostRequest{UUID: UUID, Data: data}).
				Return(v1.PostResponse{Error: err, Schema: domain.Schema{
					Inputs: []domain.Input{{Name: "email", Type: domain.EmailType}},
				}})
			return req
		}, http.StatusBadRequest},
		{http.StatusText(http.StatusInternalServerError), func(out io.Writer) *http.Request {
			req, _ := http.NewRequest(http.MethodPost, HOST, strings.NewReader("email=test@my.email"))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.SchemaKey{}, FAKE))
			service.EXPECT().
				HandlePostV1(v1.PostRequest{UUID: FAKE, Data: map[string][]string{"email": {"test@my.email"}}}).
				Return(v1.PostResponse{Error: errors.Database("", nil, "")})
			return req
		}, http.StatusInternalServerError},
		{http.StatusText(http.StatusNotFound), func(out io.Writer) *http.Request {
			req, _ := http.NewRequest(http.MethodPost, HOST, strings.NewReader("email=test@my.email"))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.SchemaKey{}, ZERO))
			service.EXPECT().
				HandlePostV1(v1.PostRequest{UUID: ZERO, Data: map[string][]string{"email": {"test@my.email"}}}).
				Return(v1.PostResponse{Error: errors.NotFound("", nil, "")})
			return req
		}, http.StatusNotFound},
		{http.StatusText(http.StatusFound), func(out io.Writer) *http.Request {
			req, _ := http.NewRequest(http.MethodPost, HOST, strings.NewReader("email=test@my.email"))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req = req.WithContext(context.WithValue(req.Context(),
				middleware.SchemaKey{}, UUID))
			service.EXPECT().
				HandlePostV1(v1.PostRequest{UUID: UUID, Data: map[string][]string{"email": {"test@my.email"}}}).
				Return(v1.PostResponse{ID: 1, Error: nil, Schema: domain.Schema{Action: HOST}})
			return req
		}, http.StatusFound},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			srv.PostV1(rw, tc.request(rw))
			assert.Equal(t, tc.code, rw.Code)
		})
	}
}
