package command

import "testing"

// TestNewCommand - test functionality of NewCommand() initializer
func TestNewCommand(t *testing.T) {
	command, err := NewCommand("test", &ModifierSet{}) // Attempt to initialize command

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
	}

	t.Logf("found command %s", command) // Log success
}
