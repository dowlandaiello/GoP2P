package node

import (
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
)

// TestNewNode - test functionality of node initialization method
func TestNewNode(t *testing.T) {
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		err = nil // Reset error

		address, err = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if err != nil { // Check second try for errors
			t.Errorf(err.Error()) // Return found error
			t.FailNow()
		}
	}

	_, err = NewNode(address, true) // Attempt to create new node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Return found error
		t.FailNow()
	}
}
