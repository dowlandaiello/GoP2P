package database

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestNewDatabase - test functionality of NewDatabase() function
func TestNewDatabase(t *testing.T) {
	address := ""                                 // Initialize address value
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil || address == "" { // Check for errors
		err = nil // Reset error

		address, err = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if err != nil { // Check second try for errors
			t.Errorf(err.Error()) // Return found error
			t.FailNow()
		}
	}

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()
	} else if err != nil { // Check that an error did indeed occur
		t.Logf(err.Error()) // Log invalid error
	}

	if err == nil {
		db, err := NewDatabase(&node, 10) // Create new database with bootstrap node, and acceptable timeout

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Fail with errors
			t.FailNow()
		}

		t.Logf("node database created successfully with bootstrap node %s", (*db.Nodes)[0].Address) // Print success
	}
}

// TestAddNode - test functionality of addNode() function
func TestAddNode(t *testing.T) {
	bootNode, err := node.NewNode("72.21.215.90", true)  // Create new node with S3 address
	secondaryNode, err := node.NewNode("1.1.1.1", false) // Create new node with Cloudflare address

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(&bootNode, 10) // Create new node database with bootstrap node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = db.AddNode(&secondaryNode) // Add Cloudflare node to database

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}

// TestRemoveNode - test functionality of removeNode() function
func TestRemoveNode(t *testing.T) {
	bootNode, err := node.NewNode("72.21.215.90", true) // Create new node with S3 address

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(&bootNode, 10) // Create new node database with bootstrap node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = db.RemoveNode("72.21.215.90") // Attempt to remove S3 node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}

func TestQueryForAddress(t *testing.T) {
	bootNode, err := node.NewNode("72.21.215.90", true) // Create new node with S3 address

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	db, err := NewDatabase(&bootNode, 10) // Create new node database with bootstrap node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	foundNodeIndex, err := db.QueryForAddress("72.21.215.90") // Search for S3 node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found node at index %s", strconv.FormatUint(uint64(foundNodeIndex), 10)) // Log success
}
