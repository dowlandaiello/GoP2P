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
		db, err := NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestnetID, 10, "test") // Create new database with bootstrap node, and acceptable timeout

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

	db, err := NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestnetID, 10, "test") // Create new node database with bootstrap node

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

	db, err := NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestnetID, 10, "test") // Create new node database with bootstrap node

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

// TestQueryForAddress - test functionality of QueryForAddress method
func TestQueryForAddress(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestnetID, 10, "test") // Create new node database with bootstrap node

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

// TestUpdateRemoteDatabase - test functionality of UpdateRemoteDatabase method
func TestUpdateRemoteDatabase(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestnetID, 10, "test") // Create new node database with bootstrap node

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: IP checking requires sudo privileges") // Log warning
	} else {
		err = db.UpdateRemoteDatabase() // Update remote instances

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log error
			t.FailNow()           // Panic
		}

		t.Logf("updated remote instances of database %s", db.NetworkAlias) // Log success
	}
}

// TestJoinDatabase - test functionality of JoinDatabase method
func TestJoinDatabase(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestnetID, 10, "test") // Create new node database with bootstrap node

	if err != nil && !strings.Contains(err.Error(), "socket") && !strings.Contains(err.Error(), "timed out") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: IP checking requires sudo privileges") // Log warning
	} else if err != nil && strings.Contains(err.Error(), "timed out") {
		t.Logf("WARNING: connection testing requires at least two nodes") // Log warning
	} else {
		err = JoinDatabase(node.Address, 3000, "GoP2P_TestNet") // Join database

		if err != nil && !strings.Contains(err.Error(), "timed out") { // Check for errors
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		} else if err != nil && strings.Contains(err.Error(), "timed out") {
			t.Logf("WARNING: connection testing requires at least two nodes") // Log warning
		} else {
			t.Logf("Joined network with alias %s", db.NetworkAlias) // Log success
		}
	}
}

// TestFetchRemoteDatabase - test functionality of FetchRemoteDatabase method
func TestFetchRemoteDatabase(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	_, err = NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestnetID, 10, "test") // Create new node database with bootstrap node

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: IP checking requires sudo privileges") // Log warning
	} else if err != nil && strings.Contains(err.Error(), "timed out") {
		t.Logf("WARNING: connection testing requires at least two nodes") // Log warning
	} else {
		fetchedDb, err := FetchRemoteDatabase(node.Address, 3000, "GoP2P_TestNet") // Fetch remote database

		if err != nil && !strings.Contains(err.Error(), "timed out") { // Check for errors
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		} else if err != nil && strings.Contains(err.Error(), "timed out") {
			t.Logf("WARNING: connection testing requires at least two nodes") // Log warning
		} else {
			t.Logf("Fetched remote database from network with alias %s", fetchedDb.NetworkAlias) // Log success
		}
	}
}

func newNodeSafe() (*node.Node, error) {
	ip, err := common.GetExtIPAddrWithoutUPnP() // Fetch IP address

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	environment, _ := environment.NewEnvironment() // Create new environment

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	localNode := node.Node{Address: ip, Reputation: 0, IsBootstrap: false, Environment: environment} // Creates new node instance with specified address

	err = localNode.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	return &localNode, nil // Return initialized node
}
