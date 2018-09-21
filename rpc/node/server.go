package rpc

import (
	"context"

	"github.com/mitsukomegumi/GoP2P/common"
	proto "github.com/mitsukomegumi/GoP2P/rpc/proto"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Server - GoP2P RPC server
type Server struct{}

// NewNode - node.NewNode RPC handler
func (server *Server) NewNode(ctx context.Context, req *proto.NewNodeRequest) (*proto.GeneralResponse, error) {
	node, err := node.NewNode(req.Address, req.IsBootstrap) // Init node

	if err != nil { // Check for errors
		return &proto.GeneralResponse{}, err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &proto.GeneralResponse{}, err // Return found error
	}

	err = node.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return &proto.GeneralResponse{}, err // Return found error
	}

	return &proto.GeneralResponse{Message: "test"}, nil // Return response
}

// StartListener - node.StartListener RPC handler
func (server *Server) StartListener(ctx context.Context, req *proto.StartListenerRequest) (*proto.GeneralResponse, error) {
	return &proto.GeneralResponse{Message: "test"}, nil // Return response
}

/* BEGIN IO HANDLERS */

// WriteToMemory - node.WriteToMemory RPC handler
func (server *Server) WriteToMemory(ctx context.Context, req *proto.MemoryRequest) (*proto.GeneralResponse, error) {
	return &proto.GeneralResponse{Message: "test"}, nil // Return response
}

// ReadFromMemory - node.ReadFromMemory RPC handler
func (server *Server) ReadFromMemory(ctx context.Context, req *proto.MemoryRequest) (*proto.GeneralResponse, error) {
	return &proto.GeneralResponse{Message: "test"}, nil // Return response
}

/* END IO HANDLERS */
