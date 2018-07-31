package database

import (
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestNewDatabase - test functionality of NewDatabase() function
func TestNewDatabase(t *testing.T) {
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		err = nil // Reset error

		address, err = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if err != nil { // Check second try for errors
			t.Errorf(err.Error()) // Return found error
			t.FailNow()
		}
	}

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()
	}

	db, err := NewDatabase(&node, 10) // Create new database with bootstrap node, and acceptable timeout

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Fail with errors
		t.FailNow()
	}

	t.Logf("node database created successfully with bootstrap node %s", (*db.Nodes)[0].Address) // Print success
}