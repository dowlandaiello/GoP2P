package environment

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/mitsukomegumi/GoP2P/common"
)

// Environment - abstract container holding variables, configurations of a certain node
type Environment struct {
	EnvironmentVariables []*Variable `json:"variables"`
}

// Variable - container holding a variable's data, and identification properties (id, type)
type Variable struct {
	VariableType           string `json:"type"`       // VariableType - type of variable (e.g. string, block, etc...)
	VariableIdentifier     string `json:"identifier"` // VariableIdentifier - id of variable (used for querying)
	VariableData           []byte `json:"data"`       // VariableData - pretty self-explanatory (usually a pointer to a struct)
	VariableSerializedData string `json:"serialized"` // VariableSerializedData - string value representation of VariableData property (used for querying)
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewEnvironment - creates new instance of environment struct with specified node value
func NewEnvironment() (*Environment, error) {
	return &Environment{EnvironmentVariables: []*Variable{{VariableType: "string", VariableIdentifier: "genesis", VariableData: []byte("Visa outage: payment chaos after card network crashes"), VariableSerializedData: "Visa outage: payment chaos after card network crashes"}}}, nil // No error occurred, return nil
}

// QueryType - Fetches latest entry into environment with matching type
func (environment *Environment) QueryType(variableType string) (*Variable, error) {
	x := len(environment.EnvironmentVariables) - 1 // Initialize iterator

	for x != -1 { // Check not out of bounds
		if environment.EnvironmentVariables[x].VariableType == variableType { // Check for matching type
			return environment.EnvironmentVariables[x], nil // Return found variable
		}

		x-- // Decrement
	}

	return &Variable{}, errors.New("no matching variable found") // No results found, return error
}

// QueryValue - searches for object with matching value
func (environment *Environment) QueryValue(value string) (*Variable, error) {
	if len(environment.EnvironmentVariables) == 0 { // Checksafe
		return &Variable{}, errors.New("found nil environment variables") // Return error
	}

	x := len(environment.EnvironmentVariables) - 1 // Initialize iterator

	for x != -1 { // Check not out of bounds
		if strings.Contains(environment.EnvironmentVariables[x].VariableSerializedData, value) { // Check for matching value
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

	serializedData, err := common.SerializeToString(variableData) // Fetch string representation

	if err != nil { // Check for errors
		return &Variable{}, err // Return found error
	}

	byteData, err := common.SerializeToBytes(variableData) // Serialize data to bytes

	if err != nil { // Check for errors
		return &Variable{}, err // Return found error
	}

	variable := Variable{VariableType: variableType, VariableIdentifier: "", VariableData: byteData, VariableSerializedData: serializedData} // Initialize variable

	serializedVariable, err := common.SerializeToBytes(variable) // Serialize variable to generate hash

	if err != nil { // Check for errors
		return &Variable{}, err // Return error
	}

	variable.VariableIdentifier = common.Sha3(serializedVariable) // Add hash to variable contents

	return &variable, nil // Return variable
}

// AddVariable - attempt to append specified variable to environment variables list
func (environment *Environment) AddVariable(variable *Variable, replaceExisting bool) error {
	if replaceExisting {
		err := environment.replaceVariable(variable)

		if err == nil {
			return nil
		}
	}

	return environment.addVariable(variable)
}

// LogEnvironment - serialize and print contents of entire environment
func (environment *Environment) LogEnvironment() error {
	marshaledVal, err := json.MarshalIndent(*environment, "", "  ") // Marshal environment

	if err != nil { // Check for errors
		return err // Return found error
	}

	fmt.Println("\n" + string(marshaledVal)) // Log marshaled val

	return nil // No error occurred, return nil
}

/*
	END EXPORTED METHODS
*/

/*
	BEGIN INTERNAL METHODS
*/

func (environment *Environment) replaceVariable(variable *Variable) error {
	foundVariable, err := environment.QueryType(variable.VariableType) // Query type

	if err != nil { // Check for errors
		return err // Return found error
	}

	(*foundVariable).VariableData = (*variable).VariableData // Set existing data to given data

	return nil // No error occurred, return nil
}

func (environment *Environment) addVariable(variable *Variable) error {
	if reflect.ValueOf(variable).IsNil() { // Check for invalid parameters
		return errors.New("invalid variable") // Return error
	}

	(*environment).EnvironmentVariables = append((*environment).EnvironmentVariables, variable) // Append value

	return nil
}

/*
	BEGIN INTERNAL METHODS
*/
