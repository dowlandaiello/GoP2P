package node

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
)

// TestNewNode - test functionality of node initialization method
func TestNewNode(t *testing.T) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("Initialized node %s", node.Address) // Log success
}

func TestStartListener(t *testing.T) {
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

	t.Logf("started listener with address %s", (*ln).Addr()) // Log success
}

func newNodeSafe() (*Node, error) {
	ip, err := common.GetExtIPAddrWithoutUpNP() // Fetch IP address

	if err != nil { // Check for errors
		return &Node{}, err // Return found error
	}

	environment, _ := environment.NewEnvironment() // Create new environment

	if err != nil { // Check for errors
		return &Node{}, err // Return found error
	}

	node := Node{Address: ip, Reputation: 0, IsBootstrap: false, Environment: environment} // Creates new node instance with specified address

	return &node, nil // Return initialized node
}
