package connection

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/types/command"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestNewConnection - test functionality of connection initialization function
func TestNewConnection(t *testing.T) {
	connection, err := generateConnection() // Create connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("created connection with source node %s", connection.InitializationNode.Address) // Log node
}

func TestAttemptConnection(t *testing.T) {
	connection, err := generateConnection() // Create connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = connection.Attempt() // Attempt connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
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
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		err = nil // Reset error

		address, err = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if err != nil { // Check second try for errors
			return nil, err // Return found error
		}
	}

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		return nil, err // Return found error
	}

	connection, err := NewConnection(&node, &node, 53, []byte("test"), "relay", []Event{}) // Attempt to initialize new connection

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return connection, nil
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
