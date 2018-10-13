package shard

import (
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// TestNewShard - test functionality of shard initializer
func TestNewShard(t *testing.T) {
	node, err := newNodeSafe() // Initialize shard node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShard(node) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("Found new shard with address %s", shard.Address) // Log new shard
}

func newNodeSafe() (*node.Node, error) {
	ip, err := common.GetExtIPAddrWithoutUpNP() // Fetch IP address

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	environment, _ := environment.NewEnvironment() // Create new environment

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	node := node.Node{Address: ip, Reputation: 0, IsBootstrap: false, Environment: environment} // Creates new node instance with specified address

	return &node, nil // Return initialized node
}
