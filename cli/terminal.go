package cli

import (
	"errors"
	"fmt"
	"reflect"
)

// Terminal - absctract container holding set of variable with values (runtime only)
type Terminal struct {
	Variables []interface{}
}

// NewTerminal - attempts to start io handler for term commands
func NewTerminal() error {
	term := Terminal{Variables: []interface{}{}}

	for {
		var input string

		fmt.Print("\n> ")
		_, err := fmt.Scanln(&input)

		if err != nil {
			return err
		}

		handleCommand(&term, string(input))
	}
}

// AddVariable - attempt to append specified variable to terminal variable list
func (term *Terminal) AddVariable(variable interface{}) error {
	if reflect.ValueOf(variable).IsNil() { // Check for nil variable
		return errors.New("nil variable found") // Return error
	}

	if len(term.Variables) == 0 { // Check for uninitialized variable array
		term.Variables = []interface{}{variable} // Initialize with variable

		return nil // No error occurred, return nil
	}

	term.Variables = append(term.Variables, variable) // Append to array

	return nil // No error occurred, return nil
}

// Attach - attempt to fetch current node configuration and save as variable
func (term *Terminal) Attach() error {
	node, err := ReadNode()

	if err != nil {
		return err
	}

	term.AddVariable(*node)

	return nil
}
