package database

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestWriteToMemory - test functionality of WriteToMemory() function
func TestWriteToMemory(t *testing.T) {
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

		dir, err := common.GetCurrentDir() // Attempt to fetch working directory

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log error
			t.FailNow()           // Panic
		}

		err = db.WriteToMemory(dir) // Attempt to write database to memory

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log error
			t.FailNow()           // Panic
		}

		t.Logf("wrote database to memory") // Log success
	}
}

// TestReadFromMemory - test functionality of ReadFromMemory() function
func TestReadFromMemory(t *testing.T) {
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

		dir, err := common.GetCurrentDir() // Attempt to fetch working directory

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log error
			t.FailNow()           // Panic
		}

		err = db.WriteToMemory(dir) // Attempt to write database to memory

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log error
			t.FailNow()           // Panic
		}

		t.Logf("wrote database to memory") // Log success

		readDb, err := ReadDatabaseFromMemory(dir) // Attempt to read database from memory

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log error
			t.FailNow()           // Panic
		}

		t.Logf("read database from memory with bootstrap node %s", (*readDb.Nodes)[0].Address) // Log success
	}
}
