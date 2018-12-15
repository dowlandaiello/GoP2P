package connection

import (
	"io"
	"strings"
	"testing"

	"github.com/dowlandaiello/GoP2P/types/command"
	"github.com/dowlandaiello/GoP2P/types/environment"
	"github.com/dowlandaiello/GoP2P/types/node"
)

// TestNewEvent - test functionality of event initializer
func TestNewEvent(t *testing.T) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()           // Panic
	} else if err != nil { // Account for special case
		t.Logf(err.Error())
	}

	resolutionValue := []byte("test")                                  // Initialize resolution value
	resolution, err := NewResolution(resolutionValue, resolutionValue) // Create resolution

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	event, err := NewEvent("push", *resolution, &command.Command{}, node, 53) // Attempt to create new event

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found event with type %s", event.EventType) // Log success
}

// TestAttemptEvent - test functionality of event initializer method
func TestAttemptEvent(t *testing.T) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()           // Panic
	}

	resolutionValue := []byte("test")                                  // Initialize resolution value
	resolution, err := NewResolution(resolutionValue, resolutionValue) // Create resolution

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	event, err := NewEvent("push", *resolution, &command.Command{}, node, 53) // Attempt to create new event

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	_, err = event.Attempt() // Attempt event

	if err != nil && !strings.Contains(err.Error(), "socket") && !strings.Contains(err.Error(), "timed out") && !strings.Contains(err.Error(), "connection refused") && err != io.EOF { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: socket actions require sudo privileges.") // Log warning
	} else if err != nil && strings.Contains(err.Error(), "connection refused") {
		t.Logf("WARNING: connection testing requires a running handler") // Log warning
	} else if err != nil && strings.Contains(err.Error(), "timed out") {
		t.Logf("WARNING: connection testing requires a running handler") // Log warning
	}
}

func newNodeSafe() (*node.Node, error) {
	ip := "1.1.1.1" // Set IP address

	environment, err := environment.NewEnvironment() // Create new environment

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	node := node.Node{Address: ip, Reputation: 0, IsBootstrap: false, Environment: environment} // Creates new node instance with specified address

	return &node, nil // Return initialized node
}
