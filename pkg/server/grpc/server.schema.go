package grpc

import (
	"context"
	"encoding/xml"
	"log"
	"strings"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kamilsk/form-api/pkg/errors"
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

// Create TODO
func (server *schemaServer) Create(ctx context.Context, req *CreateSchemaRequest) (*CreateSchemaResponse, error) {
	var data query.CreateSchema
	if err := xml.NewDecoder(strings.NewReader(req.Definition)).Decode(&data.Definition); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"trying to unmarshal XML `%s` of the schema definition",
			req.Definition)
	}
	schema, err := server.storage.CreateSchema(ctx, "", data)
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			return nil, status.Error(codes.Internal, appErr.Message())
		}
		return nil, status.Error(codes.Unknown, "trying to create the schema")
	}
	return &CreateSchemaResponse{
		Id: string(schema.ID),
		CreatedAt: &timestamp.Timestamp{
			Seconds: int64(schema.CreatedAt.Second()),
			Nanos:   int32(schema.CreatedAt.Nanosecond()),
		},
	}, nil
}

// Read TODO
func (server *schemaServer) Read(ctx context.Context, req *ReadSchemaRequest) (*ReadSchemaResponse, error) {
	log.Println("SchemaServer.Read was called")
	return &ReadSchemaResponse{}, nil
}

// Update TODO
func (server *schemaServer) Update(ctx context.Context, req *UpdateSchemaRequest) (*UpdateSchemaResponse, error) {
	log.Println("SchemaServer.Update was called")
	return &UpdateSchemaResponse{}, nil
}

// Delete TODO
func (server *schemaServer) Delete(ctx context.Context, req *DeleteSchemaRequest) (*DeleteSchemaResponse, error) {
	log.Println("SchemaServer.Delete was called")
	return &DeleteSchemaResponse{}, nil
}
