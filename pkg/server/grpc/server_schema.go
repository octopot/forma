package grpc

import (
	"bytes"
	"context"
	"encoding/xml"
	"strings"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/server/grpc/middleware"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewSchemaServer returns new instance of server API for Schema service.
func NewSchemaServer(storage ProtectedStorage) SchemaServer {
	return &schemaServer{storage: storage}
}

type schemaServer struct {
	storage ProtectedStorage
}

// Create TODO issue#173
func (server *schemaServer) Create(ctx context.Context, req *CreateSchemaRequest) (*CreateSchemaResponse, error) {
	data := query.CreateSchema{Title: req.Title}
	if err := xml.NewDecoder(strings.NewReader(req.Definition)).Decode(&data.Definition); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"trying to unmarshal XML `%s` of the schema definition",
			req.Definition)
	}
	if req.Id != "" {
		id := domain.ID(req.Id)
		data.ID = &id
	}
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	schema, err := server.storage.CreateSchema(ctx, tokenID, data)
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			if _, is = appErr.IsClientError(); is {
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
			return nil, status.Errorf(codes.Internal, "trying to create the schema %q", req.Definition)
		}
		return nil, status.Errorf(codes.Unknown, "trying to create the schema %q", req.Definition)
	}
	return &CreateSchemaResponse{
		Id: string(schema.ID),
		CreatedAt: &timestamp.Timestamp{
			Seconds: int64(schema.CreatedAt.Second()),
			Nanos:   int32(schema.CreatedAt.Nanosecond()),
		},
	}, nil
}

// Read TODO issue#173
func (server *schemaServer) Read(ctx context.Context, req *ReadSchemaRequest) (*ReadSchemaResponse, error) {
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	schema, err := server.storage.ReadSchema(ctx, tokenID, query.ReadSchema{ID: domain.ID(req.Id)})
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			if _, is = appErr.IsClientError(); is {
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
			return nil, status.Errorf(codes.Internal, "trying to read the schema %q", req.Id)
		}
		return nil, status.Errorf(codes.Unknown, "trying to read the schema %q", req.Id)
	}
	resp := ReadSchemaResponse{
		Id:    string(schema.ID),
		Title: schema.Title,
		CreatedAt: &timestamp.Timestamp{
			Seconds: int64(schema.CreatedAt.Second()),
			Nanos:   int32(schema.CreatedAt.Nanosecond()),
		},
	}
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	if encodeErr := xml.NewEncoder(buf).Encode(schema.Definition); encodeErr != nil {
		return nil, status.Errorf(codes.Internal,
			"trying to marshal definition `%#v` of the schema %q into XML",
			schema.Definition, schema.ID)
	}
	resp.Definition = buf.String()
	if schema.UpdatedAt != nil {
		resp.UpdatedAt = &timestamp.Timestamp{
			Seconds: int64(schema.UpdatedAt.Second()),
			Nanos:   int32(schema.UpdatedAt.Nanosecond()),
		}
	}
	if schema.DeletedAt != nil {
		resp.DeletedAt = &timestamp.Timestamp{
			Seconds: int64(schema.DeletedAt.Second()),
			Nanos:   int32(schema.DeletedAt.Nanosecond()),
		}
	}
	return &resp, nil
}

// Update TODO issue#173
func (server *schemaServer) Update(ctx context.Context, req *UpdateSchemaRequest) (*UpdateSchemaResponse, error) {
	data := query.UpdateSchema{ID: domain.ID(req.Id), Title: req.Title}
	if err := xml.NewDecoder(strings.NewReader(req.Definition)).Decode(&data.Definition); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"trying to unmarshal XML `%s` of the schema definition",
			req.Definition)
	}
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	schema, err := server.storage.UpdateSchema(ctx, tokenID, data)
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			if _, is = appErr.IsClientError(); is {
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
			return nil, status.Errorf(codes.Internal, "trying to update the schema %q", req.Id)
		}
		return nil, status.Errorf(codes.Unknown, "trying to update the schema %q", req.Id)
	}
	resp := UpdateSchemaResponse{}
	if schema.UpdatedAt != nil {
		resp.UpdatedAt = &timestamp.Timestamp{
			Seconds: int64(schema.UpdatedAt.Second()),
			Nanos:   int32(schema.UpdatedAt.Nanosecond()),
		}
	}
	return &resp, nil
}

// Delete TODO issue#173
func (server *schemaServer) Delete(ctx context.Context, req *DeleteSchemaRequest) (*DeleteSchemaResponse, error) {
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	schema, err := server.storage.DeleteSchema(ctx, tokenID, query.DeleteSchema{ID: domain.ID(req.Id)})
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			if _, is = appErr.IsClientError(); is {
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
			return nil, status.Errorf(codes.Internal, "trying to delete the schema %q", req.Id)
		}
		return nil, status.Errorf(codes.Unknown, "trying to delete the schema %q", req.Id)
	}
	resp := DeleteSchemaResponse{}
	if schema.DeletedAt != nil {
		resp.DeletedAt = &timestamp.Timestamp{
			Seconds: int64(schema.DeletedAt.Second()),
			Nanos:   int32(schema.DeletedAt.Nanosecond()),
		}
	}
	return &resp, nil
}
