package upnp

import (
	"log"

	upnp "github.com/NebulousLabs/go-upnp"
)

// GetGateway - get reference to current network gateway device
func GetGateway() (*upnp.IGD, error) {
	// Attempt to discover gateway device
	d, err := upnp.Discover()
	if err != nil {
		log.Fatal(err)
	}

	return d, err
}

// ForwardPort - forwards specified port on current device
func ForwardPort(portNumber int) error { // Returns error if forward failed
	GatewayDevice, err := GetGateway() // Find network gateway device

	// Check for error
	if err != nil {
		return err // Return error
	}

	// Forward specified port
	err = GatewayDevice.Forward(uint16(portNumber), "resourceforwarding")
	if err != nil {
		log.Fatal(err)
	}

	return nil // No error occurred, return nil
}
