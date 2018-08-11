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

// TestAttach - test functionality of readNode wrapper method
func TestAttach(t *testing.T) {
	node, err := NewNode() // Attempt to create new node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	node, err = AttachNode() // Attempt to read serialized node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found node with address %s", node.Address) // Log success
}
