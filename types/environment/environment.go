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
	VariableType       string      `json:"type"`       // VariableType - type of variable (e.g. string, block, etc...)
	VariableIdentifier string      `json:"identifier"` // VariableIdentifier - id of variable (used for querying)
	VariableData       interface{} `json:"data"`       // VariableData - pretty self-explanatory (usually a pointer to a struct)
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

// QueryType - Fetches latest entry into environment with matching type
func (environment *Environment) QueryType(variableType string) (*Variable, error) {
	if len(environment.EnvironmentVariables) == 0 { // Checksafe
		return &Variable{}, errors.New("found nil environment variables") // Return error
	}

	x := len(environment.EnvironmentVariables) - 1 // Initialize iterator

	for x != -1 { // Check not out of bounds
		if environment.EnvironmentVariables[x].VariableType == variableType { // Check for matching type
			return environment.EnvironmentVariables[x], nil // Return found variable
		}

		x-- // Decrement
	}

	return &Variable{}, errors.New("no matching variable found") // No results found, return error
}

// NewVariable - creates new instance of variable struct with specified types, data
func NewVariable(variableType string, variableData interface{}) (*Variable, error) {
	if variableType == "" { // Check for invalid initialization parameters
		return &Variable{}, errors.New("invalid variable initialization values") // Return error
	}

	variable := Variable{VariableType: variableType, VariableIdentifier: "", VariableData: variableData} // Initialize variable

	serializedVariable, err := common.SerializeToBytes(variable) // Serialize variable to generate hash

	if err != nil { // Check for errors
		return &Variable{}, err // Return error
	}

	variable.VariableIdentifier = common.SHA256(serializedVariable) // Add hash to variable contents

	return &variable, nil // Return variable
}

// AddVariable - attempt to append specified variable to environment variables list
func (environment *Environment) AddVariable(variable *Variable) error {
	if reflect.ValueOf(variable).IsNil() {
		return errors.New("invalid variable")
	}

	environment.EnvironmentVariables = append(environment.EnvironmentVariables, variable)

	return nil
}

/*
	END EXPORTED METHODS
*/
