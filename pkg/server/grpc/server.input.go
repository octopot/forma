package grpc

import (
	"context"
	"log"
)

// NewInputServer returns new instance of server API for Input service.
func NewInputServer() InputServer {
	return &inputServer{}
}

type inputServer struct {
}

func (*inputServer) Read(context.Context, *ReadInputRequest) (*ReadInputResponse, error) {
	log.Println("InputServer.Read was called")
	return &ReadInputResponse{}, nil
}
