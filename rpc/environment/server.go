package environment

import (
	"context"

	environmentProto "github.com/mitsukomegumi/GoP2P/rpc/proto/environment"
)

// Server - GoP2P RPC server
type Server struct{}

// NewEnvironment - environment.NewEnvironment RPC handler
func (server *Server) NewEnvironment(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {

}
