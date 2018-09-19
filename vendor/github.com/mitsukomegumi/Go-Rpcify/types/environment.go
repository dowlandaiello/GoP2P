package types

import (
	"errors"
	"reflect"
)

// Environment - stack-based container capable of holding memory-only runtime values
type Environment struct {
	Stacks []*Stack `json:"stacks"` // Fetch reference to Environment stacks
	Calls  []*Call  `json:"calls"`  // Used for storing single-reference calls (not in stacks)
}

/* BEGIN EXPORTED METHODS */

// NewEnvironment - initialize and return new instance of Environment struct
func NewEnvironment(endpoint string) *Environment {
	envStack := NewStack(endpoint) // Initialize default stack

	env := Environment{[]*Stack{envStack}, []*Call{}} // Init new environment

	return &env // Return initialized instance
}

// AddStack - attempt to append stack to environment list of stacks
func (env *Environment) AddStack(stack *Stack) error {
	if reflect.ValueOf(stack).IsNil() { // Check for nil stack
		return errors.New("nil stack") // Return found error
	}

	(*env).Stacks = append((*env).Stacks, stack) // Attempt to append to stack list

	return nil // No error occurred, return nil
}

// AddCall - attempt to append call to environment list of calls
func (env *Environment) AddCall(call *Call) error {
	if reflect.ValueOf(call).IsNil() { // Check for nil call
		return errors.New("nil call") // Return found error
	}

	(*env).Calls = append((*env).Calls, call) // Attempt to append to call list

	return nil // No error occurred, return nil
}

// SearchCallEndpoints - query for endpoint in environment calls
func (env *Environment) SearchCallEndpoints(endpoint string) (*Call, error) {
	for x := 0; x != len(env.Calls); x++ { // Check for out of bounds
		if env.Calls[x].Endpoint == endpoint { // Check for matching endpoint
			return env.Calls[x], nil // Return call
		}
	}

	return &Call{}, errors.New("call endpoint not in environment") // No stack found return nil call
}

// SearchStackEndpoints - query for endpoint in environment stacks
func (env *Environment) SearchStackEndpoints(endpoint string) (*Stack, error) {
	for x := 0; x != len(env.Stacks); x++ { // Check for out of bounds
		if env.Stacks[x].Endpoint == endpoint { // Check for matching endpoint
			return env.Stacks[x], nil // Return call
		}
	}

	return &Stack{}, errors.New("stack endpoint not in environment") // No stack was found, return nil stack
}

/* END EXPORTED METHODS */
