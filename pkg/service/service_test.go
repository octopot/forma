//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service_test -destination $PWD/pkg/service/mock_contract_test.go github.com/kamilsk/form-api/pkg/service Storage
package service_test

import (
	"errors"
	"testing"

	deep "github.com/pkg/errors"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/service"
	"github.com/kamilsk/form-api/pkg/transfer/api/v1"
	"github.com/magiconair/properties/assert"
)

const UUID domain.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestForma_HandleGetV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		dao = NewMockStorage(ctrl)
		api = service.New(dao)
	)

	tests := []struct {
		name string
		data func() (v1.GetRequest, v1.GetResponse)
	}{
		{"without error", func() (v1.GetRequest, v1.GetResponse) {
			request, response := v1.GetRequest{UUID: UUID}, v1.GetResponse{Schema: domain.Schema{}}
			dao.EXPECT().Schema(request.UUID).Return(response.Schema, nil)
			return request, response
		}},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			request, response := tc.data()
			assert.Equal(t, response, api.HandleGetV1(request))
		})
	}
}

func TestForma_HandlePostV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		dao = NewMockStorage(ctrl)
		api = service.New(dao)
	)

	tests := []struct {
		name string
		data func() (v1.PostRequest, v1.PostResponse)
	}{
		{"without error", func() (v1.PostRequest, v1.PostResponse) {
			var (
				request = v1.PostRequest{EncryptedMarker: string(UUID), UUID: UUID,
					Data: map[string][]string{"name": {"val"}},
				}
				response = v1.PostResponse{EncryptedMarker: string(UUID), ID: string(UUID), Schema: domain.Schema{
					Inputs: []domain.Input{{Name: "name", Type: domain.TextType}},
				}}
			)

			// issue #110: add cookie
			// TODO use context column
			data := map[string][]string{"name": {"val"}, "_token": {string(UUID)}}

			dao.EXPECT().Schema(request.UUID).Return(response.Schema, nil)
			dao.EXPECT().AddData(request.UUID, data).Return(response.ID, nil)
			return request, response
		}},
		{"not found error", func() (v1.PostRequest, v1.PostResponse) {
			var (
				request = v1.PostRequest{EncryptedMarker: string(UUID), UUID: UUID,
					Data: map[string][]string{"name": {"val"}},
				}
				response = v1.PostResponse{EncryptedMarker: string(UUID), Error: errors.New("not found"), Schema: domain.Schema{}}
			)
			dao.EXPECT().Schema(request.UUID).Return(response.Schema, response.Error)
			return request, response
		}},
		{"validation error", func() (v1.PostRequest, v1.PostResponse) {
			var (
				request = v1.PostRequest{EncryptedMarker: string(UUID), UUID: UUID,
					Data: map[string][]string{"email": {"test.me"}},
				}
				response = v1.PostResponse{EncryptedMarker: string(UUID), Schema: domain.Schema{
					Inputs: []domain.Input{{Name: "email", Type: domain.EmailType, Value: "test.me"}},
				}}
			)
			dao.EXPECT().Schema(request.UUID).Return(response.Schema, nil)
			_, response.Error = response.Schema.Apply(request.Data)
			return request, response
		}},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			request, response := tc.data()
			actual := api.HandlePostV1(request)
			actual.Error = deep.Cause(actual.Error)
			assert.Equal(t, response, actual)
		})
	}
}
