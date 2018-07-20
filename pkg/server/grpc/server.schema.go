package grpc

import (
	"context"
	"log"
)

// NewSchemaServer returns new instance of server API for Schema service.
func NewSchemaServer() SchemaServer {
	return &schemaServer{}
}

type schemaServer struct {
}

func (*schemaServer) Create(context.Context, *CreateSchemaRequest) (*CreateSchemaResponse, error) {
	log.Println("SchemaServer.Create was called")
	return &CreateSchemaResponse{}, nil
}

func (*schemaServer) Read(context.Context, *ReadSchemaRequest) (*ReadSchemaResponse, error) {
	log.Println("SchemaServer.Read was called")
	return &ReadSchemaResponse{}, nil
}

func (*schemaServer) Update(context.Context, *UpdateSchemaRequest) (*UpdateSchemaResponse, error) {
	log.Println("SchemaServer.Update was called")
	return &UpdateSchemaResponse{}, nil
}

func (*schemaServer) Delete(context.Context, *DeleteSchemaRequest) (*DeleteSchemaResponse, error) {
	log.Println("SchemaServer.Delete was called")
	return &DeleteSchemaResponse{}, nil
}
