package database

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestNewDatabase - test functionality of NewDatabase() function
func TestNewDatabase(t *testing.T) {
	node, err := newNodeSafe() // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()
	} else if err != nil { // Check that an error did indeed occur
		t.Logf(err.Error()) // Log invalid error
	}

	if err == nil {
		db, err := NewDatabase(node, 10) // Create new database with bootstrap node, and acceptable timeout

		if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
			t.Errorf(err.Error()) // Fail with errors
			t.FailNow()
		} else if err != nil && strings.Contains(err.Error(), "socket") {
			t.Logf("WARNING: IP checking requires sudo privileges") // Log error
		} else {
			t.Logf("node database created successfully with bootstrap node %s", (*db.Nodes)[0].Address) // Print success
		}
	}
}

// TestAddNode - test functionality of addNode() function
func TestAddNode(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(node, 10) // Create new node database with bootstrap node

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Fail with errors
		t.FailNow()
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: IP checking requires sudo privileges") // Log error
	}

	err = db.AddNode(node) // Add node to database

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Fail with errors
		t.FailNow()
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: IP checking requires sudo privileges") // Log error
	}
}

// TestRemoveNode - test functionality of removeNode() function
func TestRemoveNode(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(node, 10) // Create new node database with bootstrap node

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Fail with errors
		t.FailNow()
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: IP checking requires sudo privileges") // Log error
	} else {
		err = db.RemoveNode(node.Address) // Attempt to remove node

		if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
			t.Errorf(err.Error()) // Fail with errors
			t.FailNow()
		} else if err != nil && strings.Contains(err.Error(), "socket") {
			t.Logf("WARNING: IP checking requires sudo privileges") // Log error
		}
	}
}

func TestQueryForAddress(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(node, 10) // Create new node database with bootstrap node

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: IP checking requires sudo privileges") // Log warning
	} else {
		foundNodeIndex, err := db.QueryForAddress(node.Address) // Search for node

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		}

		t.Logf("found node at index %s", strconv.FormatUint(uint64(foundNodeIndex), 10)) // Log success
	}
}

func newNodeSafe() (*node.Node, error) {
	ip, err := common.GetExtIPAddrWithoutUpNP() // Fetch IP address

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
