package grpc

import (
	"context"
	"log"
)

// NewLogServer returns new instance of server API for Log service.
func NewLogServer(storage ProtectedStorage) LogServer {
	return &logServer{storage: storage}
}

type logServer struct {
	storage ProtectedStorage
}

// Read TODO
func (*logServer) Read(context.Context, *ReadLogsRequest) (*ReadLogsResponse, error) {
	log.Println("LogServer.Read was called")
	return &ReadLogsResponse{}, nil
}

// Listen TODO
func (*logServer) Listen(*ListenLogsRequest, Log_ListenServer) error {
	log.Println("LogServer.Listen was called")
	return nil
}
