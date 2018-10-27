package shard

import (
	"time"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Shard - container holding shard metadata
type Shard struct {
	Nodes *[]node.Node `json:"nodes"` // Nodes - primary list of nodes

	ChildNodes *[]node.Node `json:"allChildren"` // ChildNodes - list of all child nodes (recursively includes nodes in child shards, not just direct children)

	ShardRoot *Shard `json:"root"` // ShardRoot - root shard of shard tree

	Root bool `json:"isRoot"` // Root - is root

	ParentShard *Shard `json:"parent"` // ParentShard - parent of shard

	Siblings *[]*Shard `json:"siblings"` // Siblings - shard-level siblings

	ChildShards *[]Shard `json:"child shards"` // ChildShards - shards created as children of shard

	Age uint64 `json:"age"` // Age - shard age

	Origin time.Time `json:"creation time"` // Origin - time shard created

	Address string `json:"address"` // Address - addressable internet protocol ID used for shard-level communications

	ID string `json:"id"` // ID - hash of Shard contents
}

// NewShard - initialize new shard
func NewShard(initializingNode *node.Node) (*Shard, error) {
	shard := Shard{Nodes: &[]node.Node{*initializingNode}, ChildNodes: &[]node.Node{*initializingNode}, Origin: time.Now().UTC(), Address: initializingNode.Address} // Initialize shard

	serialized, err := common.SerializeToBytes(shard) // Serialize shard

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	shard.ID = common.Sha3(serialized)                                          // Set shard ID
	shard.Address, err = common.SeedAddress(initializingNode.Address, shard.ID) // Generate, set address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &shard, nil // Return initialized shard
}

// NewShardWithNodes - initialize new shard with child nodes
func NewShardWithNodes(initializingNodes *[]node.Node) (*Shard, error) {
	shard := Shard{Nodes: initializingNodes, ChildNodes: initializingNodes, Origin: time.Now().UTC(), Address: ""} // Initialize shard

	serialized, err := common.SerializeToBytes(shard) // Serialize shard

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	shard.ID = common.Sha3(serialized)                                                 // Set shard ID
	shard.Address, err = common.SeedAddress((*initializingNodes)[0].Address, shard.ID) // Generate, set address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &shard, nil // Return initialized shard
}

/*
// Shard - exponentially shard specified shard into child shards
func (shard *Shard) Shard(exponent uint) {
	if reflect.ValueOf(shard.ParentShard).IsNil() { // Check is root
		shard.Root = true       // Set root
		shard.ShardRoot = shard // Set shard rood
	}

	lastShard := shard.ShardRoot // Set last shard
	lastRoot := shard.ShardRoot  // Set root

	for x := 1; x != (len(*shard.ShardRoot.ChildNodes)/x ^ int(exponent)); x++ { // Iterate until initialized all shards

		for z := 0; z != x^int(exponent); z++ { // Initialize all siblings
			newShard, err := NewShardWithNodes(shard.ShardRoot.ChildNodes[z:])

			if err != nil { // Check for errors
				return nil, err // Return found error
			}
		}
	}
}
*/
