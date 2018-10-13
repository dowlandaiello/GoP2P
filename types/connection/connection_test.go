package connection

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/types/command"

	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestNewConnection - test functionality of connection initialization function
func TestNewConnection(t *testing.T) {
	connection, err := generateConnection() // Create connection

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: socket actions require sudo privileges.") // Log warning
	}

	t.Logf("created connection with source node %s", connection.InitializationNode.Address) // Log node
}

// TestAttemptConnection - test functionality of connection attempt() method
func TestAttemptConnection(t *testing.T) {
	connection, err := generateConnection() // Create connection

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: socket actions require sudo privileges.") // Log warning
	}

	err = connection.Attempt() // Attempt connection

	if err != nil && !strings.Contains(err.Error(), "socket") && !strings.Contains(err.Error(), "connection refused") && !strings.Contains(err.Error(), "timed out") { // Check for errors
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

// TestAttemptConnectionWithCommand - test functionality of connection attempt() method (using a command)
func TestAttemptConnectionWithCommand(t *testing.T) {
	connection, err := generateConnectionWithCommand() // Create connection

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: socket actions require sudo privileges.") // Log warning
	}

	err = connection.Attempt() // Attempt connection

	if err != nil && !strings.Contains(err.Error(), "socket") && !strings.Contains(err.Error(), "connection refused") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: socket actions require sudo privileges.") // Log warning
	} else if err != nil && strings.Contains(err.Error(), "connection refused") {
		t.Logf("WARNING: connection testing requires a running handler") // Log warning
	}
}

// TestNewResolution - test functionality of resolution initializer
func TestNewResolution(t *testing.T) {
	val := []byte("test")                      // Create temporary testing value
	resolution, err := NewResolution(val, val) // Attempt to create new resolution

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found resolution with data %s", string(resolution.ResolutionData)) // Log success
}

func generateConnection() (*Connection, error) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	connection, err := NewConnection(node, node, 53, []byte("test"), "relay", []Event{}) // Attempt to initialize new connection

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return connection, nil // No error occurred, return nil
}

func generateConnectionWithCommand() (*Connection, error) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	connection, err := NewConnection(node, node, 53, []byte("test"), "relay", *generateEventsWithCommand()) // Attempt to initialize new connection

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return connection, nil // No error occurred, return nil
}

// generateEvents - generates array of empty events (testing only)
func generateEvents() *[]Event {
	events := []Event{} // Initialize container array

	node, _ := node.NewNode("1.1.1.1", true) // Attempt to create new node

	resolutionValue := []byte("test")                                // Initialize resolution value
	resolution, _ := NewResolution(resolutionValue, resolutionValue) // Create resolution

	for x := 0; x < 15; x++ {
		event, _ := NewEvent("push", *resolution, &command.Command{}, &node, 53) // Attempt to create new event

		events = append(events, *event) // Append event to array
	}

	return &events // Return value
}

func generateEventsWithCommand() *[]Event {
	events := []Event{} // Initialize container array

	node, _ := node.NewNode("1.1.1.1", true) // Attempt to create new node

	modifierValue := "test"                                                                               // Set value for command
	command, _ := command.NewCommand("NewVariable", command.NewModifierSet("string", modifierValue, nil)) // Attempt to initialize new command

	resolutionValue := []byte("test")                                // Initialize resolution value
	resolution, _ := NewResolution(resolutionValue, resolutionValue) // Create resolution

	for x := 0; x < 15; x++ {
		event, _ := NewEvent("push", *resolution, command, &node, 53) // Attempt to create new event

		events = append(events, *event) // Append event to array
	}

	return &events // Return value
}
