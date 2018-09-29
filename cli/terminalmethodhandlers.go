package cli

import (
	"fmt"

	"github.com/mitsukomegumi/GoP2P/common"
)

// handleQueryTypeCommandTerminal - handle execution of handleQueryTypeTerminal method (wrapper)
func (term *Terminal) handleQueryTypeCommandTerminal(queryType string) {
	fmt.Println("querying type " + queryType) // Log begin

	output, err := term.handleQueryTypeTerminal(queryType) // Attempt to query for type

	if err != nil { // Check for errors
		fmt.Println("Error: " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleQueryTypeTerminal - attempt to query for specified type in environment
func (term *Terminal) handleQueryTypeTerminal(queryType string) (string, error) {
	index, err := term.QueryType(queryType) // Attempt to query for type

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	value := term.Variables[index] // Fetch value at address

	strVal, err := common.SerializeToString(value) // Serialize response

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: found variable with type " + strVal, nil // Return response
}
