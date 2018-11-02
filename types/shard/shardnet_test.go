package shard

import (
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/command"
	"github.com/mitsukomegumi/GoP2P/types/connection"
	"github.com/mitsukomegumi/GoP2P/types/node"
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

	resolution, err := connection.NewResolution([]byte("genesisFetchRequest"), "genesisFetchRequest") // Init resolution

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	command, err := command.NewCommand("QueryType", command.NewModifierSet("string", nil, nil)) // Init command

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	event, err := connection.NewEvent("fetch", *resolution, command, &node.Node{Address: (*nodeList)[0].Address}, int(3000)) // Init event

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	conn, err := connection.NewConnection(&(*nodeList)[0], &node.Node{Address: (*nodeList)[0].Address}, int(3000), []byte("dbFetchRequest"), "relay", []connection.Event{*event}) // Init connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	serializedConnection, err := common.SerializeToBytes(conn) // Serialize connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	result, err := SendBytesShardResult(serializedConnection, shard.Address, 3000) // Send bytes to shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("result: %s", result) // Log result
}

func TestSendBytesShard(t *testing.T) {
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

	err = SendBytesShard([]byte("test"), shard.Address, 3000) // Send bytes to shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("sent bytes to shard %s", shard.Address) // Log success
}
