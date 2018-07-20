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

func (*inputServer) Listen(*ListenInputRequest, Input_ListenServer) error {
	log.Println("InputServer.Listen was called")
	return nil
}
