package grpc

import (
	"bytes"
	"context"
	"encoding/xml"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.octolab.org/ecosystem/forma/internal/domain"
	"go.octolab.org/ecosystem/forma/internal/server/grpc/middleware"
	"go.octolab.org/ecosystem/forma/internal/storage/query"
)

// NewSchemaServer returns new instance of server API for Schema service.
func NewSchemaServer(storage ProtectedStorage) SchemaServer {
	return &schemaServer{storage}
}

type schemaServer struct {
	storage ProtectedStorage
}

// Create TODO issue#173
func (server *schemaServer) Create(ctx context.Context, req *CreateSchemaRequest) (*CreateSchemaResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	data := query.CreateSchema{ID: ptrToID(req.Id), Title: req.Title}
	if len(req.Definition) > 0 {
		if decodeErr := xml.NewDecoder(strings.NewReader(req.Definition)).Decode(&data.Definition); decodeErr != nil {
			return nil, status.Errorf(codes.InvalidArgument,
				"trying to unmarshal XML `%s` of the schema definition: %+v",
				req.Definition, decodeErr)
		}
	}
	schema, createErr := server.storage.CreateSchema(ctx, tokenID, data)
	if createErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", createErr)
	}
	return &CreateSchemaResponse{
		Id:        schema.ID.String(),
		CreatedAt: Timestamp(&schema.CreatedAt),
	}, nil
}

// Read TODO issue#173
func (server *schemaServer) Read(ctx context.Context, req *ReadSchemaRequest) (*ReadSchemaResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	schema, readErr := server.storage.ReadSchema(ctx, tokenID, query.ReadSchema{ID: domain.ID(req.Id)})
	if readErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", readErr)
	}
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	if encodeErr := xml.NewEncoder(buf).Encode(schema.Definition); encodeErr != nil {
		return nil, status.Errorf(codes.Internal,
			"trying to marshal definition `%#v` of the schema %q into XML: %+v",
			schema.Definition, schema.ID, encodeErr)
	}
	return &ReadSchemaResponse{
		Id:         schema.ID.String(),
		Title:      schema.Title,
		Definition: buf.String(),
		CreatedAt:  Timestamp(&schema.CreatedAt),
		UpdatedAt:  Timestamp(schema.UpdatedAt),
		DeletedAt:  Timestamp(schema.DeletedAt),
	}, nil
}

// Update TODO issue#173
func (server *schemaServer) Update(ctx context.Context, req *UpdateSchemaRequest) (*UpdateSchemaResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	data := query.UpdateSchema{ID: domain.ID(req.Id), Title: req.Title}
	if len(req.Definition) > 0 {
		if decodeErr := xml.NewDecoder(strings.NewReader(req.Definition)).Decode(&data.Definition); decodeErr != nil {
			return nil, status.Errorf(codes.InvalidArgument,
				"trying to unmarshal XML `%s` of the schema definition: %+v",
				req.Definition, decodeErr)
		}
	}
	schema, updateErr := server.storage.UpdateSchema(ctx, tokenID, data)
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", updateErr)
	}
	return &UpdateSchemaResponse{
		Id:        schema.ID.String(),
		UpdatedAt: Timestamp(schema.UpdatedAt),
	}, nil
}

// Delete TODO issue#173
func (server *schemaServer) Delete(ctx context.Context, req *DeleteSchemaRequest) (*DeleteSchemaResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	schema, deleteErr := server.storage.DeleteSchema(ctx, tokenID, query.DeleteSchema{ID: domain.ID(req.Id)})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", deleteErr)
	}
	return &DeleteSchemaResponse{
		Id:        schema.ID.String(),
		DeletedAt: Timestamp(schema.DeletedAt),
	}, nil
}
