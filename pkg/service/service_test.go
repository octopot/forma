//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service_test -destination $PWD/pkg/service/mock_storage_test.go github.com/kamilsk/form-api/pkg/service Storage
//go:generate mockgen -package service_test -destination $PWD/pkg/service/mock_handler_test.go github.com/kamilsk/form-api/pkg/service InputHandler
package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	deep "github.com/pkg/errors"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/service"
	"github.com/kamilsk/form-api/pkg/storage/types"
	"github.com/kamilsk/form-api/pkg/transfer/api/v1"
	"github.com/magiconair/properties/assert"

	_ "github.com/golang/mock/mockgen/model"
)

const UUID domain.ID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestForma_HandleGetV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		storage = NewMockStorage(ctrl)
		handler = NewMockInputHandler(ctrl)
		api     = service.New(storage, handler)
	)

	tests := []struct {
		name string
		data func() (v1.GetRequest, v1.GetResponse)
	}{
		{"without error", func() (v1.GetRequest, v1.GetResponse) {
			request, response := v1.GetRequest{ID: UUID}, v1.GetResponse{Schema: domain.Schema{}}
			storage.EXPECT().Schema(context.Background(), request.ID).Return(response.Schema, nil)
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
		storage = NewMockStorage(ctrl)
		handler = NewMockInputHandler(ctrl)
		api     = service.New(storage, handler)
	)

	tests := []struct {
		name string
		data func() (v1.PostRequest, v1.PostResponse)
	}{
		{"without error", func() (v1.PostRequest, v1.PostResponse) {
			var (
				request = v1.PostRequest{ID: UUID,
					InputData:    domain.InputData{"name": {"val"}},
					InputContext: domain.InputContext{},
				}
				response = v1.PostResponse{ID: UUID, Schema: domain.Schema{
					Inputs: []domain.Input{{Name: "name", Type: domain.TextType}},
				}}
				input = &types.Input{ID: response.ID,
					SchemaID: request.ID, Data: request.InputData, CreatedAt: time.Now(),
				}
			)

			storage.EXPECT().Schema(context.Background(), request.ID).Return(response.Schema, nil)
			handler.EXPECT().HandleInput(context.Background(), request.ID, request.InputData).Return(input, nil)
			handler.EXPECT().LogRequest(context.Background(), input, request.InputContext).Return(nil)

			return request, response
		}},
		{"not found error", func() (v1.PostRequest, v1.PostResponse) {
			var (
				request = v1.PostRequest{ID: UUID,
					InputData: domain.InputData{"name": {"val"}},
				}
				response = v1.PostResponse{Error: errors.New("not found"), Schema: domain.Schema{}}
			)
			storage.EXPECT().Schema(context.Background(), request.ID).Return(response.Schema, response.Error)
			return request, response
		}},
		{"validation error", func() (v1.PostRequest, v1.PostResponse) {
			var (
				request = v1.PostRequest{ID: UUID,
					InputData: domain.InputData{"email": {"test.me"}},
				}
				response = v1.PostResponse{Schema: domain.Schema{
					Inputs: []domain.Input{{Name: "email", Type: domain.EmailType, Value: "test.me"}},
				}}
			)
			storage.EXPECT().Schema(context.Background(), request.ID).Return(response.Schema, nil)
			_, response.Error = response.Schema.Apply(request.InputData)
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
