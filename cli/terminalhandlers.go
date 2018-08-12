package cli

import (
	"fmt"
	"strings"
)

func (term *Terminal) handleCommand(command string) {
	switch { // Iterate through possible commands
	case strings.Contains(strings.ToLower(command), "node."):
		term.handleNode(command)
	}
}

func (term *Terminal) handleNode(command string) {
	switch {
	case strings.Contains(strings.ToLower(command), "newnode"): // Account for newnode command
		term.handleNewNodeCommand()
	case strings.Contains(strings.ToLower(command), "attach"): // Account for readnode command
		term.handleAttachNodeCommand()
	}
}

func (term *Terminal) handleNewNodeCommand() {
	fmt.Println("attempting to create new node") // Log begin

	output, err := term.handleNewNode() // Attempt to initialize new node

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

func (term *Terminal) handleAttachNodeCommand() {
	fmt.Println("attempting to attach") // Log begin

	output, err := term.handleAttachNode() // Attempt to read node

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleNewNode - handle execution of NewNode() command
func (term *Terminal) handleNewNode() (string, error) {
	node, err := NewNode() // Attempt to create new node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(*node) // Add new node

	return "-- SUCCESS -- created node with address " + node.Address, nil // No error occurred, return success
}

// handleAttachNode - handle execution of ReadNode() command
func (term *Terminal) handleAttachNode() (string, error) {
	node, err := AttachNode() // Attempt to read node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(*node) // Add new node

	return "-- SUCCESS -- attached to node with address " + node.Address, nil // Log success
}
