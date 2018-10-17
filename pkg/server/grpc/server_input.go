package grpc

import (
	"context"
	"encoding/json"
	"time"

	repository "github.com/kamilsk/form-api/pkg/storage/types"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/server/grpc/middleware"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewInputServer returns new instance of server API for Input service.
func NewInputServer(storage ProtectedStorage) InputServer {
	return &inputServer{storage}
}

type inputServer struct {
	storage ProtectedStorage
}

// Read TODO issue#173
func (server *inputServer) Read(ctx context.Context, req *ReadInputRequest) (*ReadInputResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}

	var entries []*InputEntry
	push := func(input repository.Input) error {
		data, encodeErr := json.Marshal(input.Data)
		if encodeErr != nil {
			return status.Errorf(codes.Internal,
				"trying to marshal data `%#v` of the input %q into JSON: %+v",
				input.Data, input.ID, encodeErr)
		}
		entries = append(entries, &InputEntry{
			Id:       input.ID.String(),
			SchemaId: input.SchemaID.String(),
			Data:     string(data),
		})
		return nil
	}

	switch filter := req.Filter.(type) {
	case *ReadInputRequest_Condition:
		inputs, readErr := server.storage.ReadInputByFilter(ctx, tokenID, query.InputFilter{
			SchemaID: func() domain.ID {
				if filter.Condition == nil {
					return ""
				}
				return domain.ID(filter.Condition.SchemaId)
			}(),
			From: func() *time.Time {
				if filter.Condition == nil {
					return nil
				}
				if filter.Condition.CreatedAt == nil {
					return nil
				}
				return Time(filter.Condition.CreatedAt.Start)
			}(),
			To: func() *time.Time {
				if filter.Condition == nil {
					return nil
				}
				if filter.Condition.CreatedAt == nil {
					return nil
				}
				return Time(filter.Condition.CreatedAt.End)
			}(),
		})
		if readErr != nil {
			return nil, status.Errorf(codes.Internal, "error happened: %+v", readErr)
		}
		for _, input := range inputs {
			if pushErr := push(input); pushErr != nil {
				return nil, pushErr
			}
		}
	case *ReadInputRequest_Id:
		input, readErr := server.storage.ReadInputByID(ctx, tokenID, domain.ID(filter.Id))
		if readErr != nil {
			return nil, status.Errorf(codes.Internal, "error happened: %+v", readErr)
		}
		if pushErr := push(input); pushErr != nil {
			return nil, pushErr
		}
	}
	return &ReadInputResponse{Entries: entries}, nil
}
