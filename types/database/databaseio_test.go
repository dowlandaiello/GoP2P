package database

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
)

// TestWriteToMemory - test functionality of WriteToMemory() function
func TestWriteToMemory(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	if err == nil {
		db, err := NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestNetID, 10, "test") // Create new database with bootstrap node, and acceptable timeout

		if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
			t.Errorf(err.Error()) // Fail with errors
			t.FailNow()
		} else if err != nil && strings.Contains(err.Error(), "socket") {
			t.Logf("WARNING: IP checking requires sudo privileges")
		} else {
			t.Logf("node database created successfully with bootstrap node %s", (*db.Nodes)[0].Address) // Print success

			err = db.WriteToMemory(node.Environment) // Attempt to write database to memory

			if err != nil { // Check for errors
				t.Errorf(err.Error()) // Log error
				t.FailNow()           // Panic
			}

			t.Logf("wrote database to memory") // Log success
		}
	}
}

// TestReadFromMemory - test functionality of ReadFromMemory() function
func TestReadFromMemory(t *testing.T) {
	node, err := newNodeSafe() // Initialize node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else {
		db, err := NewDatabase(node, "GoP2P_TestNet", common.GoP2PTestNetID, 10, "test") // Create new database with bootstrap node, and acceptable timeout

		if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
			t.Errorf(err.Error()) // Fail with errors
			t.FailNow()
		} else if err != nil && strings.Contains(err.Error(), "socket") {
			t.Logf("WARNING: IP checking requires sudo privileges")
		} else {
			t.Logf("node database created successfully with bootstrap node %s", (*db.Nodes)[0].Address) // Print success

			err = db.WriteToMemory(node.Environment) // Attempt to write database to memory

			if err != nil { // Check for errors
				t.Errorf(err.Error()) // Log error
				t.FailNow()           // Panic
			}

			t.Logf("wrote database to memory") // Log success

			readDb, err := ReadDatabaseFromMemory(node.Environment, "GoP2P_TestNet") // Attempt to read database from memory

			if err != nil { // Check for errors
				t.Errorf(err.Error()) // Log error
				t.FailNow()           // Panic
			}

			t.Logf("read database from memory with bootstrap node %s", (*readDb.Nodes)[0].Address) // Log success
		}
	}
}
