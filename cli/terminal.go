package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
)

// Terminal - absctract container holding set of variable with values (runtime only)
type Terminal struct {
	Variables     []interface{}
	VariableTypes []string
}

// NewTerminal - attempts to start io handler for term commands
func NewTerminal() error {
	term := Terminal{Variables: []interface{}{}, VariableTypes: []string{}}

	for {
		fmt.Print("\n> ")

		reader := bufio.NewReader(os.Stdin) // Init reader

		input, err := reader.ReadString('\n') // Search for user input

		if err != nil {
			panic(err)
		}

		term.handleCommand(string(input)) // Handle specified command
	}
}

// AddVariable - attempt to append specified variable to terminal variable list
func (term *Terminal) AddVariable(variable interface{}, variableType string) error {
	if reflect.ValueOf(term).IsNil() { // Check for nil variable
		return errors.New("nil terminal found") // Return error
	}

	if len(term.Variables) == 0 { // Check for uninitialized variable array
		term.Variables = []interface{}{variable}    // Initialize with variable
		term.VariableTypes = []string{variableType} // Initialize with type

		return nil // No error occurred, return nil
	}

	term.Variables = append(term.Variables, variable)             // Append to array
	term.VariableTypes = append(term.VariableTypes, variableType) // Append to array

	return nil // No error occurred, return nil
}
