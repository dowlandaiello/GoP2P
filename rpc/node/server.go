package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/common"
	nodeProto "github.com/mitsukomegumi/GoP2P/rpc/proto/node"
	"github.com/mitsukomegumi/GoP2P/types/node"
	"github.com/mitsukomegumi/GoP2P/upnp"
)

// Server - GoP2P RPC server
type Server struct{}

// NewNode - node.NewNode RPC handler
func (server *Server) NewNode(ctx context.Context, req *nodeProto.GeneralRequest) (*nodeProto.GeneralResponse, error) {
	if req.Address == "localhost" { // Check for invalid address
		address, err := common.GetExtIPAddrWithoutUPnP() // Fetch IP

		if err == nil { // Check for errors
			req.Address = address // Set to request value
		}
	}

	node, err := node.NewNode(req.Address, req.IsBootstrap) // Init node

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	go upnp.ForwardPortSilent(3000) // Forward node port

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	err = node.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(node) // Marshal initialized variable

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	return &nodeProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// StartListener - node.StartListener RPC handler
func (server *Server) StartListener(ctx context.Context, req *nodeProto.GeneralRequest) (*nodeProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Read node from memory

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, errors.New("Node not attached") // Return found error
	}

	go node.StartListener(int(req.Port)) // Start Listener

	return &nodeProto.GeneralResponse{Message: fmt.Sprintf("\nStarted listener with host :%s", strconv.Itoa(int(req.Port)))}, nil // Return response
}

// LogNode - node.LogNode RPC handler
func (server *Server) LogNode(ctx context.Context, req *nodeProto.GeneralRequest) (*nodeProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Fetch node from working directory

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	node.LogNode() // Log node

	return &nodeProto.GeneralResponse{Message: ""}, nil // Return response
}

/* BEGIN IO HANDLERS */

// WriteToMemory - node.WriteToMemory RPC handler
func (server *Server) WriteToMemory(ctx context.Context, req *nodeProto.GeneralRequest) (*nodeProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil {
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Read node from memory

	if err != nil {
		return &nodeProto.GeneralResponse{}, errors.New("Node not attached") // Return found error
	}

	err = node.WriteToMemory(req.Path) // Write to memory

	if err != nil { // Check for errors
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	return &nodeProto.GeneralResponse{Message: "\nWrote to directory " + req.Path}, nil // Return response
}

// ReadFromMemory - node.ReadFromMemory RPC handler
func (server *Server) ReadFromMemory(ctx context.Context, req *nodeProto.GeneralRequest) (*nodeProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil {
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(req.Path) // Read node from memory

	if err != nil {
		return &nodeProto.GeneralResponse{}, errors.New("Node not found at directory " + req.Path) // Return found error
	}

	err = node.WriteToMemory(currentDir) // Attach to current directory

	if err != nil { // Check for error
		return &nodeProto.GeneralResponse{}, err // Return found error
	}

	return &nodeProto.GeneralResponse{Message: "\nRead from directory " + req.Path}, nil // Return response
}

/* END IO HANDLERS */
