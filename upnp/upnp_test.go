package upnp

import "testing"

// TestDiscoverGateway - test functionality of gateway discovery
func TestDiscoverGateway(t *testing.T) {
	_, err := GetGateway() // Attempt to fetch gateway device

	if err != nil {
		t.Errorf(err.Error()) // If error occurs, print error to testing console
	}
}

// TestForwardPort - test functionality of port forwarding
func TestForwardPort(t *testing.T) {
	err := ForwardPort(3000) // Attempt to forward port 3000

	if err != nil {
		t.Errorf(err.Error()) // If error occurs, print error to testing console
	}
}

func TestRemovePortForward(t *testing.T) {
	err := RemovePortForward(3000) // Attempt to remove forward on port 3000

	if err != nil {
		t.Errorf(err.Error()) // If error occurs, print error to testing console
	}
}
