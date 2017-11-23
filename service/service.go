package service

import "github.com/kamilsk/form-api/data/transfer/api/v1"

// New returns new instance of Form API service.
func New(dao DataLoader) *formAPI {
	return &formAPI{dao: dao}
}

type formAPI struct {
	dao DataLoader
}

// HandleGetV1 ...
func (s *formAPI) HandleGetV1(req v1.GetRequest) v1.GetResponse {
	var resp v1.GetResponse
	resp.Schema, resp.Error = s.dao.Schema(req.UUID)
	return resp
}

// HandlePostV1 ...
func (*formAPI) HandlePostV1(v1.PostRequest) v1.PostResponse {
	return v1.PostResponse{}
}
