package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mitsukomegumi/GoP2P/common"
	databaseProto "github.com/mitsukomegumi/GoP2P/rpc/proto/database"
	"github.com/mitsukomegumi/GoP2P/types/database"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Server - GoP2P RPC server
type Server struct{}

// NewDatabase - database.NewDatabase RPC handler
func (server *Server) NewDatabase(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	env, err := getLocalEnvironment(req.DataPath) // Fetch local environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	node, err := node.ReadNodeFromMemory(req.DataPath) // Attempt to read node from memory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	db, err := database.NewDatabase(node, uint(req.AcceptableTimeout)) // Create new database with bootstrap node, and acceptable timeout

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = db.WriteToMemory(env) // Write environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(db) // Marshal initialized database

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// AddNode - database.AddNode RPC handler
func (server *Server) AddNode(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	marshaledVal := ""
	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

/* BEGIN INTERNAL METHODS */

func getIP() (string, error) {
	address := ""                                 // Initialize address value
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil || address == "" { // Check for errors
		err = nil // Reset error

		address, err = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if err != nil { // Check second try for errors
			return "", err // Return found error
		}
	}

	return address, nil // Return found address
}

func getLocalEnvironment(path string) (*environment.Environment, error) {
	env, err := environment.ReadEnvironmentFromMemory(path) // Attempt to read from memory

	if err != nil { // Check for errors
		return &environment.Environment{}, err // Return found error
	}

	return env, nil // No error occurred, return found environment
}

/* END INTERNAL METHODS */
