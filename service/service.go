package service

import (
	"github.com/kamilsk/form-api/data/form"
	"github.com/kamilsk/form-api/data/transfer/api/v1"
)

// New returns new instance of Form API service.
func New(dao DataLayer) *formAPI {
	return &formAPI{dao: dao}
}

type formAPI struct {
	dao DataLayer
}

// HandleGetV1 handles `GET /api/v1/{UUID}` request.
func (s *formAPI) HandleGetV1(request v1.GetRequest) v1.GetResponse {
	var response v1.GetResponse
	response.Schema, response.Error = s.dao.Schema(request.UUID)
	return response
}

// HandlePostV1 handles `POST /api/v1/{UUID}` request.
func (s *formAPI) HandlePostV1(request v1.PostRequest) v1.PostResponse {
	var (
		response v1.PostResponse
		schema   form.Schema
		verified map[string][]string
	)
	schema, response.Error = s.dao.Schema(request.UUID)
	if response.Error != nil {
		return response
	}
	verified, response.Error = schema.Apply(request.Data)
	if response.Error != nil {
		return response
	}
	response.ID, response.Error = s.dao.AddData(request.UUID, verified)
	return response
}
