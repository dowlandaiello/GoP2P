package database

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/common"
	databaseProto "github.com/mitsukomegumi/GoP2P/internal/rpc/proto/database"
	"github.com/mitsukomegumi/GoP2P/types/database"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Server - GoP2P RPC server
type Server struct{}

/* BEGIN EXPORTED METHODS */

// NewDatabase - database.NewDatabase RPC handler
func (server *Server) NewDatabase(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	node, env, err := getLocalNodeEnvironment(currentDir) // Fetch local environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	db, err := database.NewDatabase(node, req.NetworkName, uint(req.NetworkID), uint(req.AcceptableTimeout), req.PrivateKey) // Create new database with bootstrap node, and acceptable timeout

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = db.WriteToMemory(env) // Write environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = node.WriteToMemory(currentDir) // Write environment to current dir

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
	currentDir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from current directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database, err := database.ReadDatabaseFromMemory(env, req.NetworkName) // Attempt to read database from environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	node, err := getLocalNode(currentDir) // Fetch local node

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = database.WriteToMemory(env) // Write to environment memory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = node.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database.AddNode(node) // Add node to database

	marshaledVal, err := json.Marshal((*database.Nodes)[len(*database.Nodes)-1]) // Marshal added node

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// RemoveNode - database.RemoveNode RPC handler
func (server *Server) RemoveNode(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	if req.Address == "localhost" { // Check for invalid address
		address, err := common.GetExtIPAddrWithoutUPnP()

		if err == nil { // Check for errors
			req.Address = address // Set to request value
		}
	}

	currentDir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	node, env, err := getLocalNodeEnvironment(currentDir) // Attempt to read environment from current directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database, err := database.ReadDatabaseFromMemory(env, req.NetworkName) // Attempt to read database from environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = database.RemoveNode(req.Address) // Add node to database

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = database.WriteToMemory(env) // Write to environment memory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = node.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\nRemoved node %s from database", req.Address)}, nil // Return response
}

// QueryForAddress - database.QueryForAddress RPC handler
func (server *Server) QueryForAddress(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from current directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database, err := database.ReadDatabaseFromMemory(env, req.NetworkName) // Attempt to read database from environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	nodeIndex, err := database.QueryForAddress(req.Address) // Query for address

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	foundNode := (*database.Nodes)[nodeIndex] // Fetch node at index

	marshaledVal, err := json.Marshal(foundNode) // Marshal found node

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// WriteToMemory - database.WriteToMemory RPC handler
func (server *Server) WriteToMemory(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	node, env, err := getLocalNodeEnvironment(currentDir) // Attempt to read environment from current directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database, err := database.ReadDatabaseFromMemory(env, req.NetworkName) // Attempt to read database from environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = node.WriteToMemory(req.Address) // Write to data path

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(*database) // Marshal found node

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\nWrote database %s to environment memory", string(marshaledVal))}, nil // Return response
}

// ReadFromMemory - database.ReadDatabaseFromMemory RPC handler
func (server *Server) ReadFromMemory(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	node, env, err := getLocalNodeEnvironment(req.Address) // Attempt to read environment from request path

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database, err := database.ReadDatabaseFromMemory(env, req.NetworkName) // Attempt to read database from environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = node.WriteToMemory(currentDir) // Write to current path

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(*database) // Marshal found node

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\nRead database %s from environment memory at path %s", string(marshaledVal), req.Address)}, nil // Return response
}

// UpdateRemoteDatabase - database.UpdateRemoteDatabase RPC handler
func (server *Server) UpdateRemoteDatabase(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	env, err := environment.ReadEnvironmentFromMemory(currentDir) // Attempt to read environment from request path

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database, err := database.ReadDatabaseFromMemory(env, req.NetworkName) // Attempt to read database from environment

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database.UpdateRemoteDatabase() // Update remote database instances

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\nUpdating instances of database with network alias %s and ID %s", database.NetworkAlias, strconv.Itoa(int(database.NetworkID)))}, nil // Return response
}

// JoinDatabase - database.JoinDatabase RPC handler
func (server *Server) JoinDatabase(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	err := database.JoinDatabase(req.Address, uint(req.Port), req.NetworkName) // Attempt to join network

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\nSuccessfully joined network with alias %s", req.NetworkName)}, nil // Return response
}

// FetchRemoteDatabase - database.FetchRemoteDatabase RPC handler
func (server *Server) FetchRemoteDatabase(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	db, err := database.FetchRemoteDatabase(req.Address, uint(req.Port), req.NetworkName) // Attempt to join network

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.MarshalIndent(*db, "", "  ") // Marshal found db

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// SendDatabaseMessage - database.SendDatabaseMessage RPC handler
func (server *Server) SendDatabaseMessage(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Fetch working directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	db, err := database.ReadDatabaseFromMemory(env, req.NetworkName) // Fetch database with alias

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	message, err := database.NewMessage(req.StringVals[0], uint(req.UintVal), req.StringVals[1], req.NetworkName) // Init message

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = db.SendDatabaseMessage(message, req.PrivateKey, uint(req.Port)) // Send db message

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(*message) // Marshal found node

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\nsent data %s to %s nodes", marshaledVal, strconv.Itoa(len(*db.Nodes)))}, nil // Return result message
}

// LogDatabase - database.LogDatabase RPC handler
func (server *Server) LogDatabase(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Fetch node, env from working directory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	db, err := database.ReadDatabaseFromMemory(env, req.NetworkName) // Fetch database with alias

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	db.LogDatabase() // Log database

	return &databaseProto.GeneralResponse{Message: ""}, nil // Return response
}

// FromBytes - database.FromBytes RPC handler
func (server *Server) FromBytes(ctx context.Context, req *databaseProto.GeneralRequest) (*databaseProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	env, err := environment.ReadEnvironmentFromMemory(currentDir) // Attempt to read environment from request path

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	database, err := database.FromBytes(req.ByteVal) // Fetch from byte value

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	err = database.WriteToMemory(env) // Write to environment memory

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(database) // Marshal found node

	if err != nil { // Check for errors
		return &databaseProto.GeneralResponse{}, err // Return found error
	}

	return &databaseProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

func getIP() (string, error) {
	address := ""                                 // Initialize address value
	address, err := common.GetExtIPAddrWithUPnP() // Attempt to fetch current external IP address

	if err != nil || address == "" { // Check for errors
		address, err = common.GetExtIPAddrWithoutUPnP() // Attempt to fetch address without UPnP

		if err != nil { // Check second try for errors
			return "", err // Return found error
		}
	}

	return address, nil // Return found address
}

func getLocalNodeEnvironment(path string) (*node.Node, *environment.Environment, error) {
	node, err := node.ReadNodeFromMemory(path) // Read node from memory

	if err != nil { // Check for errors
		return nil, &environment.Environment{}, err // Return found error
	}

	return node, node.Environment, nil // No error occurred, return found environment
}

func getLocalEnvironment(path string) (*environment.Environment, error) {
	node, err := node.ReadNodeFromMemory(path) // Read node from memory

	if err != nil { // Check for errors
		return &environment.Environment{}, err // Return found error
	}

	return node.Environment, nil // No error occurred, return found environment
}

func getLocalNode(path string) (*node.Node, error) {
	foundNode, err := node.ReadNodeFromMemory(path) // Read node from memory

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	return foundNode, nil // No error occurred, return found node
}

/* END INTERNAL METHODS */
