package shard

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mitsukomegumi/GoP2P/common"
	shardProto "github.com/mitsukomegumi/GoP2P/rpc/proto/shard"
	"github.com/mitsukomegumi/GoP2P/types/database"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
	"github.com/mitsukomegumi/GoP2P/types/shard"
)

// Server - GoP2P RPC server
type Server struct{}

// NewShard - shard.NewShard RPC handler
func (server *Server) NewShard(ctx context.Context, req *shardProto.GeneralRequest) (*shardProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Attempt to read node from working directory

	if err != nil { // Check for errors
		node, err = handleNoNode(req.Address) // Init node

		if err != nil { // Check for errors
			return &shardProto.GeneralResponse{}, err // Return found error
		}
	}

	db, err := database.ReadDatabaseFromMemory(node.Environment, req.NetworkName) // Read database

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	shard, err := shard.NewShard(node) // Init shard

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	err = db.AddShard(shard) // Add shard

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(*shard) // Marshal initialized variable

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	return &shardProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// NewShardWithNodes - shard.NewShardWithNodes RPC handler
func (server *Server) NewShardWithNodes(ctx context.Context, req *shardProto.GeneralRequest) (*shardProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	localNode, err := node.ReadNodeFromMemory(currentDir) // Read node from working directory

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	nodes, err := generateNodeSliceFromAddresses(req.Addresses) // Init node list

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	*nodes = append(*nodes, *localNode) // Append local node

	shard, err := shard.NewShardWithNodes(nodes) // Init shard

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(*shard) // Marshal shard

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	return &shardProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// Shard - shard.Shard RPC handler
func (server *Server) Shard(ctx context.Context, req *shardProto.GeneralRequest) (*shardProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	localNode, err := node.ReadNodeFromMemory(currentDir) // Read node from working directory

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	db, err := database.ReadDatabaseFromMemory(localNode.Environment, req.NetworkName) // Read database

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	shardIndex, err := db.QueryForShardAddress(req.Address) // Query shard by address

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	shard := (*db.Shards)[shardIndex] // Fetch shard by index

	err = shard.Shard(uint(req.Exponent)) // Shard shard

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(shard) // Marshal shard

	if err != nil { // Check for errors
		return &shardProto.GeneralResponse{}, err // Return found error
	}

	return &shardProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// generateNodeSliceFromAddress - generate nodes from node address list
func generateNodeSliceFromAddresses(addresses []string) (*[]node.Node, error) {
	nodeList := &[]node.Node{} // Init node slice buffer

	for _, address := range addresses { // Iterate through addresses
		env, _ := environment.NewEnvironment() // Init environment

		node := node.Node{Address: address, Reputation: 0, LastPingTime: time.Now().UTC(), IsBootstrap: false, Environment: env} // Init node
		*nodeList = append(*nodeList, node)                                                                                      // Append initialized node
	}

	return nodeList, nil // Return initialized node list
}

// handleNoNode - generate new node with address if no node in working directory
func handleNoNode(address string) (*node.Node, error) {
	env, err := environment.NewEnvironment() // Init environment

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	node := &node.Node{Address: address, Reputation: 0, LastPingTime: time.Now().UTC(), IsBootstrap: false, Environment: env} // Init node

	return node, nil // No error occurred, return initialized node
}
