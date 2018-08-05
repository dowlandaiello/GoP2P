package connection

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestNewConnection - test functionality of connection initialization function
func TestNewConnection(t *testing.T) {
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		err = nil // Reset error

		address, err = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if err != nil { // Check second try for errors
			t.Errorf(err.Error()) // Return found error
			t.FailNow()           // Panic
		}
	}

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()
	} else if err != nil { // Account for special case
		t.Logf(err.Error()) // Log error
	}

	connection, err := NewConnection(&node, &node, 53, []byte("test"), "relay", []Event{}) // Attempt to initialize new connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
		t.FailNow()           // Panic
	}

	t.Logf("created connection with source node %s", connection.InitializationNode.Address) // Log node
}

func TestAttemptConnection(t *testing.T) {
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		err = nil // Reset error

		address, err = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if err != nil { // Check second try for errors
			t.Errorf(err.Error()) // Return found error
			t.FailNow()           // Panic
		}
	}

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()
	} else if err != nil { // Account for special case
		t.Logf(err.Error()) // Log error
	}

	connection, err := NewConnection(&node, &node, 53, []byte("test"), "relay", []Event{}) // Attempt to initialize new connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
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
