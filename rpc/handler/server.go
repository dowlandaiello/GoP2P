package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/common"
	handlerProto "github.com/mitsukomegumi/GoP2P/rpc/proto/handler"
	"github.com/mitsukomegumi/GoP2P/types/handler"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Server - GoP2P RPC server
type Server struct{}

// StartHandler - handler.StartHandler RPC handler
func (server *Server) StartHandler(ctx context.Context, req *handlerProto.GeneralRequest) (*handlerProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &handlerProto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Attempt to read node from current directory

	if err != nil { // Check for errors
		return &handlerProto.GeneralResponse{}, errors.New("Node not attached") // Return found error
	}

	listener, err := node.StartListener(int(req.Port)) // Start Listener on specified port

	if err != nil { // Check for errors
		return &handlerProto.GeneralResponse{}, err // Return found error
	}

	go handler.StartHandler(node, listener) // Start node handler on port

	return &handlerProto.GeneralResponse{Message: fmt.Sprintf("\nStarted handler with host :%s", strconv.Itoa(int(req.Port)))}, nil // Return response
}
