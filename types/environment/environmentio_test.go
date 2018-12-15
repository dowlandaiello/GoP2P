package environment

import (
	"testing"

	"github.com/dowlandaiello/GoP2P/common"
)

// TestWriteToMemory - test functionality of WriteToMemory() method
func TestWriteToMemory(t *testing.T) {
	env, err := NewEnvironment() // Initialize new environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	currentDir, err := common.GetCurrentDir()

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = env.WriteToMemory(currentDir) // Attempt to write to memory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}

func TestReadEnvironmentFromMemory(t *testing.T) {
	env, err := NewEnvironment() // Initialize new environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	currentDir, err := common.GetCurrentDir()

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = env.WriteToMemory(currentDir) // Attempt to write to memory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	env, err = ReadEnvironmentFromMemory(currentDir) // Attempt to read environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}
