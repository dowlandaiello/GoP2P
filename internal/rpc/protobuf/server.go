package protobuf

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/internal/proto"
	protoProto "github.com/mitsukomegumi/GoP2P/internal/rpc/proto/protobuf"
)

// Server - GoP2P RPC Server
type Server struct{}

// NewProtobufGuide - proto.NewProtobufGuide RPC handler
func (server *Server) NewProtobufGuide(ctx context.Context, req *protoProto.GeneralRequest) (*protoProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	guide, err := proto.NewProtobufGuide(req.Path, req.ProtoID) // Init guide

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	err = guide.WriteToMemory(currentDir) // Write to memory at working directory

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.MarshalIndent(*guide, "", "  ") // Marshal message

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	return &protoProto.GeneralResponse{Message: fmt.Sprintf("\n%s", marshaledVal)}, nil // Return response
}

// ReadGuideFromMemory - proto.ReadGuideFromMemory RPC handler
func (server *Server) ReadGuideFromMemory(ctx context.Context, req *protoProto.GeneralRequest) (*protoProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	guide, err := proto.ReadGuideFromMemory(req.Path) // Read from path

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	err = guide.WriteToMemory(currentDir) // Write to working directory

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.MarshalIndent(*guide, "", "  ") // Marshal value

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	return &protoProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// WriteToMemory - proto.WriteToMemory RPC handler
func (server *Server) WriteToMemory(ctx context.Context, req *protoProto.GeneralRequest) (*protoProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	guide, err := proto.ReadGuideFromMemory(currentDir) // Read from path

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	err = guide.WriteToMemory(req.Path) // Write to given path

	if err != nil { // Check for errors
		return &protoProto.GeneralResponse{}, err // Return found error
	}

	return &protoProto.GeneralResponse{Message: fmt.Sprintf("\nSuccessfully wrote guide to memory at path %s", req.Path)}, nil // Return response
}
