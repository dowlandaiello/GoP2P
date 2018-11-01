package shard

import (
	"testing"
)

// TestSendBytesShardResult - test functionality of SendBytesShardResult() method
func TestSendBytesShardResult(t *testing.T) {
	nodeList, err := newNodeListSafe(8) // Init node list

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(nodeList) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	result, err := SendBytesShardResult([]byte("test"), shard.Address, 3000) // Send bytes to shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("result: %s", result) // Log result
}
