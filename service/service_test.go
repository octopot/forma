//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service_test -destination $PWD/service/mock_contract_test.go github.com/kamilsk/form-api/service Storage
package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/domain"
	"github.com/kamilsk/form-api/service"
	"github.com/kamilsk/form-api/transfer/api/v1"
	"github.com/magiconair/properties/assert"
)

const UUID domain.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestFormAPI_HandleGetV1(t *testing.T) {
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

func TestFormAPI_HandlePostV1(t *testing.T) {
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
				request  = v1.PostRequest{UUID: UUID, Data: map[string][]string{"name": {"val"}}}
				response = v1.PostResponse{ID: 1, Schema: domain.Schema{
					Inputs: []domain.Input{{Name: "name", Type: domain.TextType}},
				}}
			)
			dao.EXPECT().Schema(request.UUID).Return(response.Schema, nil)
			dao.EXPECT().AddData(request.UUID, request.Data).Return(response.ID, nil)
			return request, response
		}},
		{"not found error", func() (v1.PostRequest, v1.PostResponse) {
			var (
				request  = v1.PostRequest{UUID: UUID, Data: map[string][]string{"name": {"val"}}}
				response = v1.PostResponse{Error: errors.New("not found"), Schema: domain.Schema{}}
			)
			dao.EXPECT().Schema(request.UUID).Return(response.Schema, response.Error)
			return request, response
		}},
		{"validation error", func() (v1.PostRequest, v1.PostResponse) {
			var (
				request  = v1.PostRequest{UUID: UUID, Data: map[string][]string{"email": {"test.me"}}}
				response = v1.PostResponse{Schema: domain.Schema{
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
			assert.Equal(t, response, api.HandlePostV1(request))
		})
	}
}
