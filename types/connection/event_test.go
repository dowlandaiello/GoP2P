package connection

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

func TestNewEvent(t *testing.T) {
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
		t.FailNow()           // Panic
	} else if err != nil { // Account for special case
		t.Logf(err.Error())
	}

	resolutionValue := []byte("test")                                  // Initialize resolution value
	resolution, err := NewResolution(resolutionValue, resolutionValue) // Create resolution

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	event, err := NewEvent("push", *resolution, "", &node) // Attempt to create new event

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found event with type %s", event.EventType) // Log success
}
