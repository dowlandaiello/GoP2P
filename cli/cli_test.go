package cli

import (
	"strings"
	"testing"

	"github.com/dowlandaiello/GoP2P/common"
	"github.com/dowlandaiello/GoP2P/types/environment"
	"github.com/dowlandaiello/GoP2P/types/node"
)

// TestNewNode - test functionality of newnode wrapper method
func TestNewNode(t *testing.T) {
	node, err := NewNode() // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found node %p", node) // Log success
}

// TestAttach - test functionality of readNode wrapper method
func TestAttach(t *testing.T) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	node.WriteToMemory(currentDir) // Write node to memory

	node, err = AttachNode() // Attempt to read serialized node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found node with address %s", node.Address) // Log success
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
