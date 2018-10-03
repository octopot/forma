package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/server/grpc/middleware"
	"github.com/kamilsk/form-api/pkg/storage/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	data := query.CreateTemplate{Title: req.Title, Definition: domain.Template(req.Definition)}
	if req.Id != "" {
		id := domain.ID(req.Id)
		data.ID = &id
	}
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	template, err := server.storage.CreateTemplate(ctx, tokenID, data)
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			if _, is = appErr.IsClientError(); is {
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
			return nil, status.Errorf(codes.Internal, "trying to create the template %q", req.Definition)
		}
		return nil, status.Errorf(codes.Unknown, "trying to create the template %q", req.Definition)
	}
	return &CreateTemplateResponse{
		Id: string(template.ID),
		CreatedAt: &timestamp.Timestamp{
			Seconds: int64(template.CreatedAt.Second()),
			Nanos:   int32(template.CreatedAt.Nanosecond()),
		},
	}, nil
}

// Read TODO issue#173
func (server *templateServer) Read(ctx context.Context, req *ReadTemplateRequest) (*ReadTemplateResponse, error) {
	data := query.ReadTemplate{ID: domain.ID(req.Id)}
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	template, err := server.storage.ReadTemplate(ctx, tokenID, data)
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			if _, is = appErr.IsClientError(); is {
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
			return nil, status.Errorf(codes.Internal, "trying to read the template %q", req.Id)
		}
		return nil, status.Errorf(codes.Unknown, "trying to read the template %q", req.Id)
	}
	resp := ReadTemplateResponse{
		Id:         string(template.ID),
		Title:      template.Title,
		Definition: string(template.Definition),
		CreatedAt: &timestamp.Timestamp{
			Seconds: int64(template.CreatedAt.Second()),
			Nanos:   int32(template.CreatedAt.Nanosecond()),
		},
	}
	if template.UpdatedAt != nil {
		resp.UpdatedAt = &timestamp.Timestamp{
			Seconds: int64(template.UpdatedAt.Second()),
			Nanos:   int32(template.UpdatedAt.Nanosecond()),
		}
	}
	if template.DeletedAt != nil {
		resp.DeletedAt = &timestamp.Timestamp{
			Seconds: int64(template.DeletedAt.Second()),
			Nanos:   int32(template.DeletedAt.Nanosecond()),
		}
	}
	return &resp, nil
}

// Update TODO issue#173
func (server *templateServer) Update(ctx context.Context, req *UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	data := query.UpdateTemplate{ID: domain.ID(req.Id), Title: req.Title, Definition: domain.Template(req.Definition)}
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	template, err := server.storage.UpdateTemplate(ctx, tokenID, data)
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			if _, is = appErr.IsClientError(); is {
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
			return nil, status.Errorf(codes.Internal, "trying to update the template %q", req.Id)
		}
		return nil, status.Errorf(codes.Unknown, "trying to update the template %q", req.Id)
	}
	resp := UpdateTemplateResponse{}
	if template.UpdatedAt != nil {
		resp.UpdatedAt = &timestamp.Timestamp{
			Seconds: int64(template.UpdatedAt.Second()),
			Nanos:   int32(template.UpdatedAt.Nanosecond()),
		}
	}
	return &resp, nil
}

// Delete TODO issue#173
func (server *templateServer) Delete(ctx context.Context, req *DeleteTemplateRequest) (*DeleteTemplateResponse, error) {
	data := query.DeleteTemplate{ID: domain.ID(req.Id)}
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	template, err := server.storage.DeleteTemplate(ctx, tokenID, data)
	if err != nil {
		if appErr, is := err.(errors.ApplicationError); is {
			if _, is = appErr.IsClientError(); is {
				return nil, status.Error(codes.InvalidArgument, appErr.Message())
			}
			return nil, status.Errorf(codes.Internal, "trying to delete the template %q", req.Id)
		}
		return nil, status.Errorf(codes.Unknown, "trying to delete the template %q", req.Id)
	}
	resp := DeleteTemplateResponse{}
	if template.DeletedAt != nil {
		resp.DeletedAt = &timestamp.Timestamp{
			Seconds: int64(template.DeletedAt.Second()),
			Nanos:   int32(template.DeletedAt.Nanosecond()),
		}
	}
	return &resp, nil
}
