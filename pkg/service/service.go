package service

import (
	"context"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"github.com/kamilsk/form-api/pkg/transfer/api/v1"
)

// New returns a new instance of the Forma service.
func New(storage Storage, handler InputHandler) *Forma {
	return &Forma{storage: storage, handler: handler}
}

// Forma is the primary application service.
type Forma struct {
	storage Storage
	handler InputHandler
}

// HandleGetV1 handles an input request.
func (service *Forma) HandleGetV1(request v1.GetRequest) v1.GetResponse {
	var response v1.GetResponse
	response.Schema, response.Error = service.storage.Schema(context.Background(), request.ID)
	return response
}

// HandlePostV1 handles an input request.
func (service *Forma) HandlePostV1(request v1.PostRequest) v1.PostResponse {
	var (
		response v1.PostResponse
		verified domain.InputData
	)

	response.Schema, response.Error = service.storage.Schema(context.Background(), request.ID)
	if response.Error != nil {
		return response
	}
	verified, response.Error = response.Schema.Apply(request.InputData)
	if response.Error != nil {
		response.Error = errors.Validation(errors.InvalidFormDataMessage, response.Error,
			"trying to add data for schema %q", request.ID)
		return response
	}

	var input *query.Input
	input, response.Error = service.handler.HandleInput(context.Background(), request.ID, verified)
	if response.Error != nil {
		return response
	}
	response.ID = input.ID
	go service.handler.LogRequest(context.Background(), input, request.InputContext)
	return response
}
