package cli

import "testing"

// TestNewNode - test functionality of newnode wrapper method
func TestNewNode(t *testing.T) {
	node, err := NewNode() // Attempt to create new node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found node with address %s", node.Address) // Log success
}

// TestReadNode - test functionality of readNode wrapper method
func TestReadNode(t *testing.T) {
	node, err := NewNode() // Attempt to create new node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	node, err = ReadNode() // Attempt to read serialized node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found node with address %s", node.Address) // Log success
}
