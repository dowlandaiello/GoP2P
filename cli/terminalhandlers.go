package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mitsukomegumi/GoP2P/types/node"
	"github.com/mitsukomegumi/GoP2P/upnp"
)

func (term *Terminal) handleCommand(command string) {
	switch { // Iterate through possible commands
	case strings.Contains(strings.ToLower(command), "upnp."): // Account for UpNP methods
		term.handleUpNP(command)
	case strings.Contains(strings.ToLower(command), "node."): // Account for node methods
		term.handleNode(command)
	}
}

/*
	BEGIN METHOD ROUTING
*/

func (term *Terminal) handleUpNP(command string) {
	switch {
	case strings.Contains(strings.ToLower(command), "forwardport"): // Account for forwardport command
		intVal, _ := strconv.Atoi(strings.Split(strings.Split(command, "(")[1], ")")[0]) // Attempt to fetch port from command

		term.handleForwardPortCommand(intVal) // Forward port
	case strings.Contains(strings.ToLower(command), "removeportforward"): // Account for removeportforward command
		intVal, _ := strconv.Atoi(strings.Split(strings.Split(command, "(")[1], ")")[0]) // Attempt to fetch port from command

		term.handleRemoveForwardPortCommand(intVal) // Remove port forwarding
	}
}

func (term *Terminal) handleNode(command string) {
	switch {
	case strings.Contains(strings.ToLower(command), "newnode"): // Account for newnode command
		term.handleNewNodeCommand()
	case strings.Contains(strings.ToLower(command), "attach"): // Account for readnode command
		term.handleAttachNodeCommand()
	case strings.Contains(strings.ToLower(command), "startlistener"):
		intVal, _ := strconv.Atoi(strings.Split(strings.Split(command, "(")[1], ")")[0]) // Attempt to fetch port from command

		term.handleStartListenerCommand(intVal) // Start listener command execution
	}
}

/*
	END METHOD ROUTING
*/

/*
	BEGIN UpNP METHODS
*/

func (term *Terminal) handleForwardPortCommand(portNumber int) {
	fmt.Println("attempting to forward port") // Log begin

	output, err := term.handleForwardPort(portNumber) // Attempt to forward port

	if err != nil { // Check for errors
		fmt.Println(err.Error()) // log found error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleForwardPort - handle execution of forwardport method
func (term *Terminal) handleForwardPort(portNumber int) (string, error) {
	fmt.Println("attempting to remove port forwarding") // Log begin

	err := upnp.ForwardPort(uint(portNumber)) // Attempt to forward port

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- port " + strconv.Itoa(portNumber) + " forwarded successfully", nil // Return success
}

func (term *Terminal) handleRemoveForwardPortCommand(portNumber int) {
	output, err := term.handleRemoveForwardPort(portNumber) // Attempt to remove port forwarding

	if err != nil { // Check for errors
		fmt.Println(err.Error()) // log found error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleForwardPort - handle execution of removeportforward method
func (term *Terminal) handleRemoveForwardPort(portNumber int) (string, error) {
	err := upnp.RemovePortForward(uint(portNumber)) // Attempt to remove port forwarding

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- forwarding on port " + strconv.Itoa(portNumber) + " removed successfully", nil // Return success
}

/*
	END UpNP METHODS
*/

/*
	BEGIN NODE METHODS
*/

// handleNewNodeCommand - handle execution of newnode method
func (term *Terminal) handleNewNodeCommand() {
	fmt.Println("attempting to create new node") // Log begin

	output, err := term.handleNewNode() // Attempt to initialize new node

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleAttachNodeCommand - attempt to read node at current working directory
func (term *Terminal) handleAttachNodeCommand() {
	fmt.Println("attempting to attach") // Log begin

	output, err := term.handleAttachNode() // Attempt to read node

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleStartListenerCommand - attempt to start listener on attached node
func (term *Terminal) handleStartListenerCommand(port int) {
	fmt.Println("attempting to start listener") // Log begin

	output, err := term.handleStartListener(port) // Attempt to read node

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

	term.AddVariable(*node, "Node") // Add new node

	return "-- SUCCESS -- created node with address " + node.Address, nil // No error occurred, return success
}

// handleAttachNode - handle execution of ReadNode() command
func (term *Terminal) handleAttachNode() (string, error) {
	node, err := AttachNode() // Attempt to read node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(*node, "Node") // Add new node

	return "-- SUCCESS -- attached to node with address " + node.Address, nil // Log success
}

func (term *Terminal) handleStartListener(port int) (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Node" { // Verify element is node
			foundNode = term.Variables[x].(node.Node) // Set to value
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	ln, err := foundNode.StartListener(port) // Attempt to start listener

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = term.AddVariable(*ln, "Listener") // Attempt to save

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- started listener on port " + strconv.Itoa(port) + " with address " + foundNode.Address, nil
}

/*
	END NODE METHODS
*/
