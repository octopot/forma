package grpc

import (
	"context"
	"log"
)

// NewTemplateServer returns new instance of server API for Template service.
func NewTemplateServer(storage ProtectedStorage) TemplateServer {
	return &templateServer{storage: storage}
}

type templateServer struct {
	storage ProtectedStorage
}

// Create TODO
func (*templateServer) Create(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	log.Println("TemplateServer.Create was called")
	return &CreateTemplateResponse{}, nil
}

// Read TODO
func (*templateServer) Read(context.Context, *ReadTemplateRequest) (*ReadTemplateResponse, error) {
	log.Println("TemplateServer.Read was called")
	return &ReadTemplateResponse{}, nil
}

// Update TODO
func (*templateServer) Update(context.Context, *UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	log.Println("TemplateServer.Update was called")
	return &UpdateTemplateResponse{}, nil
}

// Delete TODO
func (*templateServer) Delete(context.Context, *DeleteTemplateRequest) (*DeleteTemplateResponse, error) {
	log.Println("TemplateServer.Delete was called")
	return &DeleteTemplateResponse{}, nil
}
