package protobuf

import (
	"context"
	"encoding/json"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/internal/proto"
	protoProto "github.com/mitsukomegumi/GoP2P/internal/rpc/proto/proto"
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

	return &protoProto.GeneralResponse{Message: string(marshaledVal)}, nil // No error occurred, return marshaled guide
}
