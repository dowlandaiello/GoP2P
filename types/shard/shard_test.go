package shard

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

/*
	BEGIN EXPORTED METHODS
*/

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

	t.Logf("Initialized new shard with address %s", shard.Address) // Log new shard
}

// TestNewShardWithNodes - test functionality of shard initializer
func TestNewShardWithNodes(t *testing.T) {
	localNode, err := newNodeSafe() // Initialize shard node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(&[]node.Node{*localNode, *localNode}) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("Initialized new shard with address %s", shard.Address) // Log new shard
}

// TestShardShard - test functionality of exponential sharding
func TestShardShard(t *testing.T) {
	nodeList, err := newNodeListSafe(4) // Initialize shard node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(nodeList) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = shard.Shard(2) // Shard shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = shard.LogShard() // Log shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	serialized, err := common.SerializeToString(*shard) // Serialize shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	writeTest(serialized, "Subsharding") // Write serialized
}

func TestSerializeShard(t *testing.T) {
	nodeList, err := newNodeListSafe(4) // Initialize shard node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(nodeList) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = shard.Shard(2) // Shard shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	serialized, err := common.SerializeToBytes(shard) // Serialize shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("Serialized shard %s", string(serialized)) // Log success
}

// TestQueryForAddress - test functionality of shard address querying
func TestQueryForAddress(t *testing.T) {
	nodeList, err := newNodeListSafe(2) // Initialize shard node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(nodeList) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	index, err := shard.QueryForAddress((*nodeList)[0].Address) // Query for address

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found node index: %s", strconv.Itoa(int(index))) // Log success
}

// TestLogShard - test functionality of shard logging
func TestLogShard(t *testing.T) {
	nodeList, err := newNodeListSafe(2) // Initialize shard node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(nodeList) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = shard.LogShard() // Log shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}

// TestCalculateQuadraticExponent - test functionality of quadratic exponent calculator
func TestCalculateQuadraticExponent(t *testing.T) {
	t.Logf(fmt.Sprintf("%f", CalculateQuadraticExponent(4))) // Log output
}

/*
	END EXPORTED METHODS
*/

/*
	BEGIN INTERNAL METHODS
*/

func newNodeSafe() (*node.Node, error) {
	ip := "1.1.1.1" // Cloudflare IP

	environment, err := environment.NewEnvironment() // Create new environment

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	node := node.Node{Address: ip, Reputation: 0, IsBootstrap: false, Environment: environment} // Creates new node instance with specified address

	return &node, nil // Return initialized node
}

func newNodeListSafe(nodeCount int) (*[]node.Node, error) {
	nodeList := &[]node.Node{} // Init node list

	localNode, err := newNodeSafe() // Init local node

	if err != nil { // Check for errors
		return &[]node.Node{}, err // Return found error
	}

	for x := 0; x != nodeCount; x++ { // Iterate until nodeCount reached
		*nodeList = append(*nodeList, *localNode) // Append node
	}

	return nodeList, nil // No error occurred, return nil
}

func writeTest(data string, testName string) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	file, err := os.Create(currentDir + filepath.FromSlash("/test") + testName + ".json") // Attempt to create file at path

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	writer := bufio.NewWriter(file) // Init writer

	_, err = writer.WriteString(data) // Write data

	if err != nil { // Check for errors
		panic(err) // Panic
	}
}

/*
	END INTERNAL METHODS
*/
