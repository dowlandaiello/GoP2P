package database

import (
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

func TestNewDatabase(t *testing.T) {
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Return found error
	}

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Return found error
	}

	db, err := NewDatabase(&node, 10)

	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("node database created successfully with bootstrap node %s", (*db.Nodes)[0].Address)
}
