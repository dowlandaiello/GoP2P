package common

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mitsukomegumi/GoP2P/common"
	commonProto "github.com/mitsukomegumi/GoP2P/rpc/proto/common"
)

// Server - GoP2P RPC server
type Server struct{}

// ParseStringMethodCall - common.ParseStringMethodCall RPC handler
func (server *Server) ParseStringMethodCall(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	receiver, methodName, params, err := common.ParseStringMethodCall(req.Input) // Parse string method call

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(params) // Marshal initialized database

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("receiver: %s, methodname: %s, params: %s", receiver, methodName, marshaledVal)}, nil // Return response
}
