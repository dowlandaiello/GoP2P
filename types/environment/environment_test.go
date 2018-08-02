package environment

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

func TestNewEnvironment(t *testing.T) {
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

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()
	} else if err != nil { // Account for special case
		t.Logf(err.Error())
	}

	env, err := NewEnvironment(&node) // Initialize new environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	serializedEnv, err := common.SerializeToBytes(env) // Serialize environment

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	envID := common.SHA256(serializedEnv) // Get hash value of environment

	t.Logf("created environment with ID %s", envID) // Log success
}

// TestNewVariable - test functionality of variable initialization function
func TestNewVariable(t *testing.T) {
	variable, err := NewVariable("string", "test") // Create new string variable

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("created variable %s", variable) // Log success
}
