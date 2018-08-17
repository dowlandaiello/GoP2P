package cli

import (
	"fmt"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/upnp"
)

/*
	BEGIN UpNP METHODS
*/

func (term *Terminal) handleForwardPortCommand(command string, portNumber int) {
	fmt.Println("attempting to forward port " + strconv.Itoa(portNumber)) // Log begin

	output, err := term.handleForwardPort(command, portNumber) // Attempt to forward port

	if err != nil { // Check for errors
		fmt.Println(err.Error()) // log found error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleForwardPort - handle execution of forwardport method
func (term *Terminal) handleForwardPort(command string, portNumber int) (string, error) {
	err := upnp.ForwardPort(uint(portNumber)) // Attempt to forward port

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	if hasVariableSet(command) {
		term.handleOutputVariable(command, "-- SUCCESS -- port "+strconv.Itoa(portNumber)+" forwarded successfully", "string")
	}

	return "-- SUCCESS -- port " + strconv.Itoa(portNumber) + " forwarded successfully", nil // Return success
}

func (term *Terminal) handleRemoveForwardPortCommand(command string, portNumber int) {
	fmt.Println("attempting remove forwarding on port " + strconv.Itoa(portNumber)) // Log begin

	output, err := term.handleRemoveForwardPort(command, portNumber) // Attempt to remove port forwarding

	if err != nil { // Check for errors
		fmt.Println(err.Error()) // log found error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleForwardPort - handle execution of removeportforward method
func (term *Terminal) handleRemoveForwardPort(command string, portNumber int) (string, error) {
	err := upnp.RemovePortForward(uint(portNumber)) // Attempt to remove port forwarding

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	if hasVariableSet(command) {
		term.handleOutputVariable(command, "-- SUCCESS -- forwarding on port "+strconv.Itoa(portNumber)+" removed successfully", "string")
	}

	return "-- SUCCESS -- forwarding on port " + strconv.Itoa(portNumber) + " removed successfully", nil // Return success
}

/*
	END UpNP METHODS
*/
