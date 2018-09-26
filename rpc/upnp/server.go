package upnp

import (
	"context"
	"encoding/json"
	"fmt"

	upnpProto "github.com/mitsukomegumi/GoP2P/rpc/proto/upnp"
	"github.com/mitsukomegumi/GoP2P/upnp"
)

// Server - GOP2P RPC server
type Server struct{}

// GetGateway - upnp.GetGateway RPC handler
func (server *Server) GetGateway(ctx context.Context, req *upnpProto.GeneralRequest) (*upnpProto.GeneralResponse, error) {
	gateway, err := upnp.GetGateway() // Attempt to fetch gateway device

	if err != nil { // Check for errors
		return &upnpProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(*gateway) // Marshal initialized variable

	if err != nil { // Check for errors
		return &upnpProto.GeneralResponse{}, err // Return found error
	}

	return &upnpProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // No error occurred, return nil
}
