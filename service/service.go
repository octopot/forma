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
func (*formAPI) HandleGetV1(v1.GetRequest) v1.GetResponse {
	return v1.GetResponse{}
}

// HandlePostV1 ...
func (*formAPI) HandlePostV1(v1.PostRequest) v1.PostResponse {
	return v1.PostResponse{}
}
