package cli

import (
	"fmt"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/upnp"
)

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
