package environment

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/mitsukomegumi/GoP2P/common"
	environmentProto "github.com/mitsukomegumi/GoP2P/rpc/proto/environment"
	"github.com/mitsukomegumi/GoP2P/types/environment"
)

// Server - GoP2P RPC server
type Server struct{}

// NewEnvironment - environment.NewEnvironment RPC handler
func (server *Server) NewEnvironment(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := environment.ReadEnvironmentFromMemory(currentDir) // Attempt to read environment from memory

	if err != nil { // Check for errors
		env, err = environment.NewEnvironment() // Init environment

		if err != nil { // Check for errors
			return &environmentProto.GeneralResponse{}, err // Return found error
		}
	}

	env.WriteToMemory(currentDir) // Save for persistency

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\nInitialized environment %v", env)}, nil // Return response
}

// QueryType - environment.QueryType RPC handler
func (server *Server) QueryType(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := environment.ReadEnvironmentFromMemory(currentDir) // Attempt to read environment from memory

	if err != nil { // Check for errors
		env, err = environment.NewEnvironment() // Init environment

		if err != nil { // Check for errors
			return &environmentProto.GeneralResponse{}, err // Return found error
		}

		env.WriteToMemory(currentDir) // Save for persistency
	}

	foundVariable, err := env.QueryType(req.VariableType) // Query type

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\nFound variable %v", foundVariable)}, nil // No error occurred, return output
}

// QueryValue - environment.QueryValue RPC handler
func (server *Server) QueryValue(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := environment.ReadEnvironmentFromMemory(currentDir) // Attempt to read environment from memory

	if err != nil { // Check for errors
		env, err = environment.NewEnvironment() // Init environment

		if err != nil { // Check for errors
			return &environmentProto.GeneralResponse{}, err // Return found error
		}

		env.WriteToMemory(currentDir) // Save for persistency
	}

	foundVariable, err := env.QueryValue(req.Value) // Query type

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\nFound variable %v", foundVariable)}, nil // No error occurred, return output
}

// NewVariable - environment.NewVariable RPC handler
func (server *Server) NewVariable(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := environment.ReadEnvironmentFromMemory(currentDir) // Attempt to read environment from memory

	if err != nil { // Check for errors
		env, err = environment.NewEnvironment() // Init environment

		if err != nil { // Check for errors
			return &environmentProto.GeneralResponse{}, err // Return found error
		}

		env.WriteToMemory(currentDir) // Save for persistency
	}

	data, err := ioutil.ReadFile(req.Path) // Read file

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	variable, err := environment.NewVariable(req.VariableType, string(data)) // Query type

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\nInitialized variable %v", variable)}, nil // No error occurred, return output
}

// AddVariable - environment.AddVariable RPC handler
func (server *Server) AddVariable(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := environment.ReadEnvironmentFromMemory(currentDir) // Attempt to read environment from memory

	if err != nil { // Check for errors
		env, err = environment.NewEnvironment() // Init environment

		if err != nil { // Check for errors
			return &environmentProto.GeneralResponse{}, err // Return found error
		}

		env.WriteToMemory(currentDir) // Save for persistency
	}

	variable := &environment.Variable{} // Init buffer

	err = common.ReadGob(req.Path, environment.Variable{}) // Read file

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	err = env.AddVariable(variable, true) // Attempt to add variable

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\nAdded variable %v to Environment", *variable)}, nil // No error occurred, return output
}
