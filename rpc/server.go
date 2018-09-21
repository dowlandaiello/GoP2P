package rpc

import (
	"context"

	proto "github.com/mitsukomegumi/GoP2P/rpc/proto"
)

// Server - gop2p RPC server
type Server struct{}

// NewNode - node.NewNode RPC handler
func (server *Server) NewNode(ctx context.Context, req *proto.NewNodeRequest) (*proto.GeneralResponse, error) {
	return &proto.GeneralResponse{Message: "test"}, nil // Return response
}

// StartListener - node.StartListener RPC handler
func (server *Server) StartListener(ctx context.Context, req *proto.StartListenerRequest) (*proto.GeneralResponse, error) {
	return &proto.GeneralResponse{Message: "test"}, nil // Return response
}

/* BEGIN IO HANDLERS */

// WriteToMemory - node.WriteToMemory RPC handler
func (server *Server) WriteToMemory(ctx context.Context, req *proto.NewNodeRequest) (*proto.GeneralResponse, error) {
	return &proto.GeneralResponse{Message: "test"}, nil // Return response
}

// ReadFromMemory - node.ReadFromMemory RPC handler
func (server *Server) ReadFromMemory(ctx context.Context, req *proto.ReadFromMemoryRequest) (*proto.GeneralResponse, error) {
	return &proto.GeneralResponse{Message: "test"}, nil // Return response
}

/* END IO HANDLERS */
