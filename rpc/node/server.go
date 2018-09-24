package rpc

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/common"
	proto "github.com/mitsukomegumi/GoP2P/rpc/proto"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Server - GoP2P RPC server
type Server struct{}

// NewNode - node.NewNode RPC handler
func (server *Server) NewNode(ctx context.Context, req *proto.GeneralRequest) (*proto.GeneralResponse, error) {
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

	return &proto.GeneralResponse{Message: fmt.Sprintf("\nInitialized node %v", node)}, nil // Return response
}

// StartListener - node.StartListener RPC handler
func (server *Server) StartListener(ctx context.Context, req *proto.GeneralRequest) (*proto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &proto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Read node from memory

	if err != nil { // Check for errors
		return &proto.GeneralResponse{}, errors.New("Node not attached") // Return found error
	}

	go node.StartListener(int(req.Port)) // Start Listener

	return &proto.GeneralResponse{Message: fmt.Sprintf("Started listener with host :%s", strconv.Itoa(int(req.Port)))}, nil // Return response
}

/* BEGIN IO HANDLERS */

// WriteToMemory - node.WriteToMemory RPC handler
func (server *Server) WriteToMemory(ctx context.Context, req *proto.GeneralRequest) (*proto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil {
		return &proto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Read node from memory

	if err != nil {
		return &proto.GeneralResponse{}, errors.New("Node not attached") // Return found error
	}

	err = node.WriteToMemory(req.Path) // Write to memory

	if err != nil { // Check for errors
		return &proto.GeneralResponse{}, err // Return found error
	}

	return &proto.GeneralResponse{Message: "Wrote to directory " + req.Path}, nil // Return response
}

// ReadFromMemory - node.ReadFromMemory RPC handler
func (server *Server) ReadFromMemory(ctx context.Context, req *proto.GeneralRequest) (*proto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil {
		return &proto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(req.Path) // Read node from memory

	if err != nil {
		return &proto.GeneralResponse{}, errors.New("Node not fount at directory " + req.Path) // Return found error
	}

	err = node.WriteToMemory(currentDir) // Attach to current directory

	if err != nil { // Check for error
		return &proto.GeneralResponse{}, err // Return found error
	}

	return &proto.GeneralResponse{Message: "Read from directory " + req.Path}, nil // Return response
}

/* END IO HANDLERS */
