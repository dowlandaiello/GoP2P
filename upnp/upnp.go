package upnp

import (
	"fmt"
	"strconv"
	"time"

	upnp "github.com/NebulousLabs/go-upnp"
	"github.com/briandowns/spinner"
)

/*
	BEGIN EXPORTED METHODS:
*/

// GetGateway - get reference to current network gateway device
func GetGateway() (*upnp.IGD, error) { // Returns error if forward failed, returns gateway device is succeeded
	d, err := upnp.Discover() // Attempt to discover gateway device

	if err != nil { // Check for errors
		return nil, err // Return error
	}

	return d, nil // Return gateway
}

// ForwardPortSilent - forwards specified port on current device without log output
func ForwardPortSilent(portNumber uint) error { // Returns error if forward failed
	GatewayDevice, err := GetGateway() // Find network gateway device

	if err != nil { // Check for errors
		return err // Return error
	}

	err = GatewayDevice.Forward(uint16(portNumber), "resourceforwarding") // Attempt to forward

	if err != nil { // Check if error occurred
		return err // Return error
	}

	return nil // No error occurred, return nil
}

// ForwardPort - forwards specified port on current device
func ForwardPort(portNumber uint) error { // Returns error if forward failed
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond) // Init loading indicator

	s.Prefix = "   "                                     // Add line spacing
	s.Suffix = " attempting to discover network gateway" // Add log message

	s.Start() // Start loading indicator

	GatewayDevice, err := GetGateway() // Find network gateway device

	if err != nil { // Check for errors
		s.Stop() // Stop loading indicator

		return err // Return error
	}

	s.Stop() // Stop loading indicator

	fmt.Println("\nfound gateway") // Log found gateway

	s.Suffix = " attempting to forward port" // Add log message

	s.Restart() // Start loading indicator

	err = GatewayDevice.Forward(uint16(portNumber), "resourceforwarding") // Attempt to forward

	if err != nil { // Check if error occurred
		s.Stop()   // Stop loading indicator
		return err // Return error
	}

	s.Stop() // Stop loading indicator

	fmt.Println("\nsuccessfully forwarded port " + strconv.Itoa(int(portNumber))) // Log success

	return nil // No error occurred, return nil
}

// RemovePortForward - removes all forwarding for specified port
func RemovePortForward(portNumber uint) error { // Returns error if removal failed
	GatewayDevice, err := GetGateway() // Find network gateway device

	if err != nil { // Check for errors
		return err // Return error
	}

	err = GatewayDevice.Clear(uint16(portNumber)) // Remove specified port forwarding

	if err != nil { // Check for errors
		return err // Return error
	}

	return nil // No error occurred, return nil
}

/*
	END EXPORTED METHODS
*/
