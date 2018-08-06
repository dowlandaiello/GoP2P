package handler

import (
	"strings"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

func TestStartHandler(t *testing.T) {
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		err = nil // Reset error

		address, err = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if err != nil { // Check second try for errors
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		}
	}

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil { // Account for special case
		t.Logf(err.Error()) // Log found error
	}

	ln, err := node.StartListener(3000) // Start listener

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	go func() {
		err = StartHandler(&node, ln) // Attempt to start handler

		if err != nil { // Check for error
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		}
	}()

	t.Logf("started handler with listener address %s", (*ln).Addr()) // Log success
}
