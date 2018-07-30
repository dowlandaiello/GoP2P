package upnp

import (
	upnp "github.com/NebulousLabs/go-upnp"
)

// GetGateway - get reference to current network gateway device
func GetGateway() (*upnp.IGD, error) { // Returns error if forward failed, returns gateway device is succeeded
	// Attempt to discover gateway device
	d, err := upnp.Discover()
	if err != nil {
		return nil, err // Return error
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
		return err // Return error
	}

	return nil // No error occurred, return nil
}

// RemovePortForward - removes all forwarding for specified port
func RemovePortForward(portNumber int) error { // Returns error if removal failed
	GatewayDevice, err := GetGateway() // Find network gateway device

	// Check for error
	if err != nil {
		return err // Return error
	}

	// Remove specified port forwarding
	err = GatewayDevice.Clear(uint16(portNumber))
	if err != nil {
		return err // Return error
	}

	return nil // No error occurred, return nil
}
