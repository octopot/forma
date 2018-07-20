package grpc

import (
	"context"
	"log"
)

// NewLogServer returns new instance of server API for Log service.
func NewLogServer() LogServer {
	return &logServer{}
}

type logServer struct {
}

func (*logServer) Read(context.Context, *ReadLogRequest) (*ReadLogResponse, error) {
	log.Println("LogServer.Read was called")
	return &ReadLogResponse{}, nil
}

func (*logServer) Listen(*ListenLogRequest, Log_ListenServer) error {
	log.Println("LogServer.Listen was called")
	return nil
}
