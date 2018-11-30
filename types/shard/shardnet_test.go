package shard

import (
	"testing"
)

// TestSendBytesShardResult - test functionality of SendBytesShardResult() method
func TestSendBytesShardResult(t *testing.T) {
	nodeList, err := newNodeListSafe(2) // Init node list

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(nodeList) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	result, err := SendBytesShardResult([]byte("test"), shard.Address, 443) // Send bytes to shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("result: %s", result) // Log result
}

func TestSendBytesShard(t *testing.T) {
	nodeList, err := newNodeListSafe(2) // Init node list

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(nodeList) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = SendBytesShard([]byte("test"), shard.Address, 443) // Send bytes to shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("sent bytes to shard %s", shard.Address) // Log success
}
