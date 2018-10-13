package common

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/common"
	commonProto "github.com/mitsukomegumi/GoP2P/rpc/proto/common"
)

// Server - GoP2P RPC server
type Server struct{}

// SeedAddress - common.SeedAddress RPC handler
func (server *Server) SeedAddress(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	seededAddress, err := common.SeedAddress(req.Input, req.SecondInput) // Seed address

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: seededAddress}, nil // Return response
}

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

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\nreceiver: %s, methodname: %s, params: %s", receiver, methodName, string(marshaledVal))}, nil // Return response
}

// ParseStringParams - common.ParseStringParams RPC handler
func (server *Server) ParseStringParams(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	params, err := common.ParseStringParams(req.Input) // Parse string params

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(params) // Marshal initialized database

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// ConvertStringToReflectValues - common.ConvertStringToReflectValues RPC handler
func (server *Server) ConvertStringToReflectValues(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	reflectValues := common.ConvertStringToReflectValues(req.Inputs) // Parse string params

	marshaledVal, err := json.Marshal(reflectValues) // Marshal initialized database

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// StringStripReceiverCall - common.StringStripReceiverCall RPC handler
func (server *Server) StringStripReceiverCall(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	stripped := common.StringStripReceiverCall(req.Input) // Parse string params

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", stripped)}, nil // Return response
}

// StringStripParentheses - common.StringStripParentheses RPC handler
func (server *Server) StringStripParentheses(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	stripped := common.StringStripParentheses(req.Input) // Strip parentheses

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", stripped)}, nil // Return response
}

// StringFetchCallReceiver - common.StringFetchCallReceiver RPC handler
func (server *Server) StringFetchCallReceiver(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	receiver := common.StringFetchCallReceiver(req.Input) // Fetch receiver

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", receiver)}, nil // Return response
}

// CheckAddress - common.CheckAddress RPC handler
func (server *Server) CheckAddress(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	err := common.CheckAddress(req.Input) // Attempt to check the address of S3

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\nAddress %s verified", req.Input)}, nil // Return response
}

// GetExtIPAddrWithUpNP - common.GetExtIPAddrWithUpNP RPC handler
func (server *Server) GetExtIPAddrWithUpNP(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	ip, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch external IP address

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", ip)}, nil // Return response
}

// GetExtIPAddrWithoutUpNP - common.GetExtIPAddrWithoutUpNP RPC handler
func (server *Server) GetExtIPAddrWithoutUpNP(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	ip, err := common.GetExtIPAddrWithoutUpNP() // Attempt to fetch external IP address

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", ip)}, nil // Return response
}

// GetCurrentTime - common.GetCurrentTime RPC handler
func (server *Server) GetCurrentTime(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	time := common.GetCurrentTime() // Fetch current time

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", time.String())}, nil // Return response
}

// GetCurrentDir - common.GetCurrentDir RPC handler
func (server *Server) GetCurrentDir(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	dir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", dir)}, nil // Return response
}

// SHA256 - common.SHA256 RPC handler
func (server *Server) SHA256(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	hash := common.SHA256(req.ByteInput) // Hash

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\n%s", hash)}, nil // Return response
}

// SendBytes - common.SendBytes RPC handler
func (server *Server) SendBytes(ctx context.Context, req *commonProto.GeneralRequest) (*commonProto.GeneralResponse, error) {
	err := common.SendBytes(req.ByteInput, req.Input) // Send bytes

	if err != nil { // Check for errors
		return &commonProto.GeneralResponse{}, err // Return found error
	}

	return &commonProto.GeneralResponse{Message: fmt.Sprintf("\nSent %s bytes to address %s successfully", strconv.Itoa(len(req.ByteInput)), req.Input)}, nil // Return response
}
