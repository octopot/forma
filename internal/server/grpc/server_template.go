package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.octolab.org/ecosystem/forma/internal/domain"
	"go.octolab.org/ecosystem/forma/internal/server/grpc/middleware"
	"go.octolab.org/ecosystem/forma/internal/storage/query"
)

// NewTemplateServer returns new instance of server API for Template service.
func NewTemplateServer(storage ProtectedStorage) TemplateServer {
	return &templateServer{storage}
}

type templateServer struct {
	storage ProtectedStorage
}

// Create TODO issue#173
func (server *templateServer) Create(ctx context.Context, req *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	template, createErr := server.storage.CreateTemplate(ctx, tokenID, query.CreateTemplate{
		ID:         ptrToID(req.Id),
		Title:      req.Title,
		Definition: domain.Template(req.Definition),
	})
	if createErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", createErr)
	}
	return &CreateTemplateResponse{
		Id:        template.ID.String(),
		CreatedAt: Timestamp(&template.CreatedAt),
	}, nil
}

// Read TODO issue#173
func (server *templateServer) Read(ctx context.Context, req *ReadTemplateRequest) (*ReadTemplateResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	template, readErr := server.storage.ReadTemplate(ctx, tokenID, query.ReadTemplate{ID: domain.ID(req.Id)})
	if readErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", readErr)
	}
	return &ReadTemplateResponse{
		Id:         template.ID.String(),
		Title:      template.Title,
		Definition: string(template.Definition),
		CreatedAt:  Timestamp(&template.CreatedAt),
		UpdatedAt:  Timestamp(template.UpdatedAt),
		DeletedAt:  Timestamp(template.DeletedAt),
	}, nil
}

// Update TODO issue#173
func (server *templateServer) Update(ctx context.Context, req *UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	template, updateErr := server.storage.UpdateTemplate(ctx, tokenID, query.UpdateTemplate{
		ID:         domain.ID(req.Id),
		Title:      req.Title,
		Definition: domain.Template(req.Definition),
	})
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", updateErr)
	}
	return &UpdateTemplateResponse{
		Id:        template.ID.String(),
		UpdatedAt: Timestamp(template.UpdatedAt),
	}, nil
}

// Delete TODO issue#173
func (server *templateServer) Delete(ctx context.Context, req *DeleteTemplateRequest) (*DeleteTemplateResponse, error) {
	tokenID, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	template, deleteErr := server.storage.DeleteTemplate(ctx, tokenID, query.DeleteTemplate{ID: domain.ID(req.Id)})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "error happened: %+v", deleteErr)
	}
	return &DeleteTemplateResponse{
		Id:        template.ID.String(),
		DeletedAt: Timestamp(template.DeletedAt),
	}, nil
}
