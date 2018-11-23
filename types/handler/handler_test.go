package handler

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestStartHandler - test functionality of StartHandler
func TestStartHandler(t *testing.T) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil { // Account for special case
		t.Logf(err.Error()) // Log found error
	}

	ln, err := node.StartListener(3000) // Start listener

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	go func() {
		err = StartHandler(node, ln) // Attempt to start handler

		if err != nil { // Check for error
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		}
	}()

	t.Logf("started handler with listener address %s", (*ln).Addr()) // Log success
}

// TestStartProtobufHandler - test functionality of StartProtobufHandler
func TestStartProtobufHandler(t *testing.T) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil { // Account for special case
		t.Logf(err.Error()) // Log found error
	}

	ln, err := node.StartListener(3000) // Start listener

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	go func() {
		err = StartProtobufHandler(testHandler, "test", ln) // Attempt to start handler

		if err != nil { // Check for error
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		}
	}()
}

// testHandler - test handler
func testHandler(b []byte) error {
	fmt.Println(string(b)) // Log input

	return nil // No error occurred
}

func newNodeSafe() (*node.Node, error) {
	ip, err := common.GetExtIPAddrWithoutUPnP() // Fetch IP address

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	environment, _ := environment.NewEnvironment() // Create new environment

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	node := node.Node{Address: ip, Reputation: 0, IsBootstrap: false, Environment: environment} // Creates new node instance with specified address

	return &node, nil // Return initialized node
}
