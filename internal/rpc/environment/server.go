package environment

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mitsukomegumi/GoP2P/common"
	environmentProto "github.com/mitsukomegumi/GoP2P/internal/rpc/proto/environment"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Server - GoP2P RPC server
type Server struct{}

// NewEnvironment - environment.NewEnvironment RPC handler
func (server *Server) NewEnvironment(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from memory

	if err != nil { // Check for errors
		env, err = environment.NewEnvironment() // Init environment

		if err != nil { // Check for errors
			return &environmentProto.GeneralResponse{}, err // Return found error
		}
	}

	env.WriteToMemory(currentDir) // Save for persistency

	marshaledVal, err := json.Marshal(*env) // Marshal initialized variable

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // Return response
}

// QueryType - environment.QueryType RPC handler
func (server *Server) QueryType(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from memory

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

	marshaledVal, err := json.Marshal(foundVariable) // Marshal found value

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // No error occurred, return output
}

// QueryValue - environment.QueryValue RPC handler
func (server *Server) QueryValue(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from memory

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

	marshaledVal, err := json.Marshal(foundVariable) // Marshal found value

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // No error occurred, return output
}

// NewVariable - environment.NewVariable RPC handler
func (server *Server) NewVariable(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from memory

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

	marshaledVal, err := json.Marshal(variable) // Marshal initialized variable

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // No error occurred, return output
}

// AddVariable - environment.AddVariable RPC handler
func (server *Server) AddVariable(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from memory

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

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\nAdded variable %v to Environment", *variable)}, nil // No error occurred, return output
}

// WriteToMemory - environment.WriteToMemory RPC handler
func (server *Server) WriteToMemory(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from memory

	if err != nil { // Check for errors
		env, err = environment.NewEnvironment() // Init environment

		if err != nil { // Check for errors
			return &environmentProto.GeneralResponse{}, err // Return found error
		}
	}

	err = env.WriteToMemory(req.Path) // Save for persistency

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\nWrote environment %s to path %s", env, req.Path)}, nil // No error occurred, return output
}

// ReadFromMemory - environment.ReadEnvironmentFromMemory RPC handler
func (server *Server) ReadFromMemory(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	env, err := getLocalEnvironment(req.Path) // Attempt to read environment from memory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	err = env.WriteToMemory(req.Path) // Save for persistency

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	marshaledVal, err := json.Marshal(*env) // Marshal initialized variable

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: fmt.Sprintf("\n%s", string(marshaledVal))}, nil // No error occurred, return output
}

// LogEnvironment - environment.LogEnvironment RPC handler
func (server *Server) LogEnvironment(ctx context.Context, req *environmentProto.GeneralRequest) (*environmentProto.GeneralResponse, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	env, err := getLocalEnvironment(currentDir) // Attempt to read environment from memory

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	err = env.LogEnvironment() // Log environment

	if err != nil { // Check for errors
		return &environmentProto.GeneralResponse{}, err // Return found error
	}

	return &environmentProto.GeneralResponse{Message: ""}, nil // No error occurred, return output
}

/* BEGIN INTERNAL METHODS */

func getLocalEnvironment(path string) (*environment.Environment, error) {
	node, err := node.ReadNodeFromMemory(path) // Read node from path

	if err != nil { // Check for errors
		return &environment.Environment{}, err // Return found error
	}

	return node.Environment, nil // No error occurred, return environment
}

/* END INTERNAL METHODS */
