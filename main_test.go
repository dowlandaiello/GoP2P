package main

import (
	"strings"
	"testing"

	"github.com/dowlandaiello/GoP2P/common"
	"github.com/dowlandaiello/GoP2P/types/node"
)

// TestMain - test functionality of main() method
func TestMain(t *testing.T) {
	currentDir, err := common.GetCurrentDir() // Get working directory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	node, err := node.NewNode("1.1.1.1", false) // Init node

	if err != nil && !strings.Contains(err.Error(), "not permitted") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "not permitted") { // Check for perm errors
		t.Logf("WARNING: node testing requires sudo privileges") // Log sudo required
	} else {
		err = node.WriteToMemory(currentDir) // Write to working directory

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		}

		go main() // :shrug:
	}
}

// TestStartRPCServer - test functionality of startRPCServer() method
func TestStartRPCServer(t *testing.T) {
	startRPCServer() // Start RPC server
}

// TestStartNode - test functionality of startNode() method
func TestStartNode(t *testing.T) {
	currentDir, err := common.GetCurrentDir() // Get working directory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	node, err := node.NewNode("1.1.1.1", false) // Init node

	if err != nil && !strings.Contains(err.Error(), "not permitted") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "not permitted") { // Check for perm errors
		t.Logf("WARNING: node testing requires sudo privileges") // Log sudo required
	} else {
		err = node.WriteToMemory(currentDir) // Write to working directory

		if err != nil { // Check for errors
			t.Errorf(err.Error()) // Log found error
			t.FailNow()           // Panic
		}

		go startNode() // Start local node
	}
}
