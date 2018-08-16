// +build integration

//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package main -destination $PWD/mock_storage_test.go github.com/kamilsk/form-api/service Storage
package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/server/router/chi"
	"github.com/kamilsk/form-api/pkg/service"
	"github.com/kamilsk/form-api/pkg/transfer/encoding"
	"github.com/stretchr/testify/assert"
)

const (
	HOST  = "http://form-api.dev/"
	APIv1 = "api/v1"
	FAKE  = domain.ID("41ca5e09-3ce2-0094-b108-3ecc257c6fa4")
	ZERO  = domain.ID("00000000-0000-4000-8000-000000000000")
	UUID  = domain.ID("41ca5e09-3ce2-4094-b108-3ecc257c6fa4")
)

func TestAPI_GetV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		storage = NewMockStorage(ctrl)
	)

	handler := chi.NewRouter(config.ServerConfig{BaseURL: HOST, TemplateDir: "static/templates"}, service.New(storage))
	srv := httptest.NewServer(handler)
	defer srv.Close()

	{
		var (
			schema = domain.Schema{
				Title:        "Email subscription",
				Action:       "https://kamil.samigullin.info/",
				Method:       "post",
				EncodingType: "application/x-www-form-urlencoded",
				Inputs: []domain.Input{
					{
						Name:      "email",
						Type:      domain.EmailType,
						Title:     "Email",
						MaxLength: 64,
						Required:  true,
					},
				},
			}
		)

		tests := []struct {
			name    string
			request *http.Request
			golden  string
		}{
			{"get schema as HTML", func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, join(srv.URL, APIv1, UUID.String()), nil)
				if err != nil {
					panic(err)
				}
				req.Header.Set("Accept", encoding.HTML)
				return req
			}(), "./transfer/encoding/fixtures/email_subscription.html.golden"},
			{"get schema as JSON", func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, join(srv.URL, APIv1, UUID.String()), nil)
				if err != nil {
					panic(err)
				}
				req.Header.Set("Accept", encoding.JSON)
				return req
			}(), "./transfer/encoding/fixtures/email_subscription.json.golden"},
			{"get schema as XML", func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, join(srv.URL, APIv1, UUID.String()), nil)
				if err != nil {
					panic(err)
				}
				req.Header.Set("Accept", encoding.XML)
				return req
			}(), "./transfer/encoding/fixtures/email_subscription.xml.golden"},
			{"get schema as text", func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, join(srv.URL, APIv1, UUID.String()), nil)
				if err != nil {
					panic(err)
				}
				req.Header.Set("Accept", encoding.TEXT)
				return req
			}(), "./transfer/encoding/fixtures/email_subscription.yaml.golden"},
		}
		storage.EXPECT().Schema(UUID).Times(len(tests)).Return(schema, nil)

		for _, test := range tests {
			tc := test
			t.Run(test.name, func(t *testing.T) {
				resp, err := http.DefaultClient.Do(tc.request)
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, resp.StatusCode)

				expected, err := ioutil.ReadFile(tc.golden)
				assert.NoError(t, err)
				obtained, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.NoError(t, resp.Body.Close())

				// this is because server contains domain logic
				assert.NotEqual(t, string(expected), string(obtained))
			})
		}
	}

	{
		tests := []struct {
			name    string
			request *http.Request
			code    int
		}{
			{http.StatusText(http.StatusBadRequest), func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, join(srv.URL, APIv1, FAKE.String()), nil)
				if err != nil {
					panic(err)
				}
				return req
			}(), http.StatusBadRequest},
			{http.StatusText(http.StatusNotAcceptable), func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, join(srv.URL, APIv1, UUID.String()), nil)
				if err != nil {
					panic(err)
				}
				req.Header.Set("Accept", "application/toml")
				return req
			}(), http.StatusNotAcceptable},
			{http.StatusText(http.StatusNotFound), func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, join(srv.URL, APIv1, ZERO.String()), nil)
				if err != nil {
					panic(err)
				}
				return req
			}(), http.StatusNotFound},
		}
		storage.EXPECT().Schema(ZERO).Times(1).Return(domain.Schema{}, errors.NotFound("", nil, ""))

		for _, test := range tests {
			tc := test
			t.Run(test.name, func(t *testing.T) {
				resp, err := http.DefaultClient.Do(tc.request)
				assert.NoError(t, err)
				assert.Equal(t, tc.code, resp.StatusCode)
			})
		}
	}
}

func TestAPI_PostV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}

func join(base string, paths ...string) string {
	u, err := url.Parse(base)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(append([]string{u.Path}, paths...)...)
	return u.String()
}
