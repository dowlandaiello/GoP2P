package rpc

import (
	"context"

	proto "github.com/mitsukomegumi/gop2p/rpc/proto"
)

// Server - gop2p RPC server
type Server struct{}

// NewNode - node.NewNode RPC handler
func (server *Server) NewNode(ctx context.Context, req *proto.NewNodeRequest) {

}
