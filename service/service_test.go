//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package service_test -destination $PWD/service/mock_contract_test.go github.com/kamilsk/form-api/service DataLayer
package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/form"
	"github.com/kamilsk/form-api/data/transfer/api/v1"
	"github.com/kamilsk/form-api/service"
	"github.com/magiconair/properties/assert"
)

const UUID data.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestFormAPI_HandleGetV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := v1.GetRequest{UUID: UUID}
	resp := v1.GetResponse{Error: nil, Schema: form.Schema{}}
	dao := NewMockDataLayer(ctrl)
	dao.EXPECT().Schema(UUID).Return(resp.Schema, resp.Error)
	api := service.New(dao)

	assert.Equal(t, resp, api.HandleGetV1(req))
}

func TestFormAPI_HandlePostV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	{
		req := v1.PostRequest{UUID: UUID, Data: map[string][]string{"name1": {"val1"}}}
		resp := v1.PostResponse{Error: nil, ID: 1, Schema: form.Schema{Inputs: []form.Input{
			{Name: "name1", Type: form.TextType}}}}
		dao := NewMockDataLayer(ctrl)
		dao.EXPECT().Schema(UUID).Return(resp.Schema, resp.Error)
		dao.EXPECT().AddData(UUID, req.Data).Return(resp.ID, resp.Error)
		api := service.New(dao)

		assert.Equal(t, resp, api.HandlePostV1(req))
	}

	{
		err := errors.New("not found")
		req := v1.PostRequest{UUID: UUID, Data: map[string][]string{"name1": {"val1"}}}
		resp := v1.PostResponse{Error: err, ID: 0, Schema: form.Schema{}}
		dao := NewMockDataLayer(ctrl)
		dao.EXPECT().Schema(UUID).Return(form.Schema{}, err)
		api := service.New(dao)

		assert.Equal(t, resp, api.HandlePostV1(req))
	}

	{
		req := v1.PostRequest{UUID: UUID, Data: map[string][]string{"name1": {"val1"}}}
		resp := v1.PostResponse{ID: 0, Schema: form.Schema{Inputs: []form.Input{{Name: "name1", Type: form.EmailType}}}}
		_, resp.Error = resp.Schema.Apply(req.Data)
		dao := NewMockDataLayer(ctrl)
		dao.EXPECT().Schema(UUID).Return(resp.Schema, nil)
		api := service.New(dao)

		assert.Equal(t, resp, api.HandlePostV1(req))
	}
}
