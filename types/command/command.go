package command

import (
	"errors"
	"reflect"
)

// Command - absctract container holding command values
type Command struct {
	Command string `json:"command"`

	Modifiers *Modifiers `json:"modifiers"`
}

// Modifiers - abstract containers holding specific parameters for a command
type Modifiers struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewCommand - attempt to initialize new instance of command struct with specified command, modifiers.
func NewCommand(command string, modifiers *Modifiers) (*Command, error) {
	if command == "" { // Check for nil command
		return &Command{}, errors.New("invalid command") // Return found error
	} else if reflect.ValueOf(modifiers).IsNil() { // Check for nil modifier
		return &Command{}, errors.New("invalid modifier") // Return found error
	}

	return &Command{Command: command, Modifiers: modifiers}, nil // Return instance
}
