package cli

import (
	"fmt"
	"strings"
)

func handleCommand(term *Terminal, command string) {
	switch { // Iterate through possible commands
	case strings.Contains(strings.ToLower(command), "newnode"): // Account for newnode command
		fmt.Println("attempting to create new node") // Log begin

		output, err := handleNewNode(term) // Attempt to initialize new node

		if err != nil { // Check for errors
			fmt.Println("-- ERROR -- " + err.Error()) // Log error
		} else {
			fmt.Println(output) // Log success
		}
	case strings.Contains(strings.ToLower(command), "attach"): // Account for readnode command
		fmt.Println("attempting to attach") // Log begin

		output, err := handleNewNode(term) // Attempt to read node

		if err != nil { // Check for errors
			fmt.Println("-- ERROR -- " + err.Error()) // Log error
		} else {
			fmt.Println(output) // Log success
		}
	}
}

// handleNewNode - handle execution of NewNode() command
func handleNewNode(term *Terminal) (string, error) {
	node, err := NewNode() // Attempt to create new node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(*node) // Add new node

	return "-- SUCCESS -- created node with address " + node.Address, nil // No error occurred, return success
}

// handleReadNode - handle execution of ReadNode() command
func handleAttachNode(term *Terminal) (string, error) {
	node, err := AttachNode() // Attempt to read node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(*node) // Add new node

	return "-- SUCCESS -- attached to node with address " + node.Address, nil // Log success
}
