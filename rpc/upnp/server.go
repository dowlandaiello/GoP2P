package upnp

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

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

// ForwardPortSilent - upnp.ForwardPortSilent RPC handler
func (server *Server) ForwardPortSilent(ctx context.Context, req *upnpProto.GeneralRequest) (*upnpProto.GeneralResponse, error) {
	err := upnp.ForwardPortSilent(uint(req.PortNumber)) // Forward specified port

	if err != nil { // Check for errors
		return &upnpProto.GeneralResponse{}, err // Return found error
	}

	return &upnpProto.GeneralResponse{Message: fmt.Sprintf("\nForwarded port %s", strconv.Itoa(int(req.PortNumber)))}, nil // No error occurred, return nil
}

// ForwardPort - upnp.ForwardPort RPC handler
func (server *Server) ForwardPort(ctx context.Context, req *upnpProto.GeneralRequest) (*upnpProto.GeneralResponse, error) {
	err := upnp.ForwardPort(uint(req.PortNumber)) // Forward specified port

	if err != nil { // Check for errors
		return &upnpProto.GeneralResponse{}, err // Return found error
	}

	return &upnpProto.GeneralResponse{Message: fmt.Sprintf("\nForwarded port %s", strconv.Itoa(int(req.PortNumber)))}, nil // No error occurred, return nil
}

// RemoveForwarding - upnp.RemovePortForward RPC handler
func (server *Server) RemoveForwarding(ctx context.Context, req *upnpProto.GeneralRequest) (*upnpProto.GeneralResponse, error) {
	err := upnp.RemovePortForward(uint(req.PortNumber)) // Remove forwarding on specified port

	if err != nil { // Check for errors
		return &upnpProto.GeneralResponse{}, err // Return found error
	}

	return &upnpProto.GeneralResponse{Message: fmt.Sprintf("\nRemoved forwarding %s", strconv.Itoa(int(req.PortNumber)))}, nil // No error occurred, return nil
}
