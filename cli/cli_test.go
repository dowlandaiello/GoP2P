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
