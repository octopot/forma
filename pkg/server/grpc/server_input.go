package grpc

import (
	"context"
	"log"
)

// NewInputServer returns new instance of server API for Input service.
func NewInputServer(storage ProtectedStorage) InputServer {
	return &inputServer{storage}
}

type inputServer struct {
	storage ProtectedStorage
}

// Read TODO issue#173
func (*inputServer) Read(context.Context, *ReadInputRequest) (*ReadInputResponse, error) {
	log.Println("InputServer.Read was called")
	return &ReadInputResponse{}, nil
}
