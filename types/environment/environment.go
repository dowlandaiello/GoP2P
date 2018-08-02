package environment

import (
	"errors"
	"reflect"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Environment - abstract container holding variables, configurations of a certain node
type Environment struct {
	EnvironmentVariables []*Variable `json:"variables"`
	EnvironmentNode      *node.Node  `json:"node"`
}

// Variable - container holding a variable's data (pointer), and identification properties (id, type)
type Variable struct {
	VariableType       string       `json:"type"`       // VariableType - type of variable (e.g. string, block, etc...)
	VariableIdentifier string       `json:"identifier"` // VariableIdentifier - id of variable (used for querying)
	VariableData       *interface{} `json:"data"`       // VariableData - pretty self-explanatory (usually a pointer to a struct)
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewEnvironment - creates new instance of environment struct with specified node value
func NewEnvironment(node *node.Node) (*Environment, error) {
	if reflect.ValueOf(node).IsNil() { // Check that node is not nil
		return nil, errors.New("invalid node") // Return error if true
	}

	return &Environment{EnvironmentVariables: []*Variable{}, EnvironmentNode: node}, nil // No error occurred, return nil
}

// NewVariable - creates new instance of variable struct with specified types, data
func NewVariable(variableType string, variableData *interface{}) (*Variable, error) {
	if reflect.ValueOf(variableData).IsNil() || variableType == "" {
		return &Variable{}, errors.New("invalid variable initialization values")
	}

	variable := Variable{VariableType: variableType, VariableIdentifier: "", VariableData: variableData}

	serializedVariable, err := common.SerializeToBytes(variable)

	if err != nil {
		return &Variable{}, err
	}

	variable.VariableIdentifier = common.SHA256(serializedVariable)

	return &variable, nil
}

/*
	END EXPORTED METHODS
*/
