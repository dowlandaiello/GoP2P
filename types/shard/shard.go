package shard

import (
	"time"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Shard - container holding shard metadata
type Shard struct {
	Nodes *[]node.Node `json:"nodes"` // Nodes - primary list of nodes

	ChildShards *[]Shard `json:"child shards"` // ChildShards - shards created as children of shard

	Age uint64 `json:"age"` // Age - shard age

	Origin time.Time `json:"creation time"` // Origin - time shard created

	Address string `json:"address"` // Address - addressable internet protocol ID used for shard-level communications

	ID string `json:"id"` // ID - hash of Shard contents
}

// NewShard - initialize new shard
func NewShard(initializingNode *node.Node) (*Shard, error) {
	shard := Shard{Nodes: &[]node.Node{*initializingNode}, Origin: time.Now().UTC(), Address: initializingNode.Address} // Initialize shard

	serialized, err := common.SerializeToBytes(shard) // Serialize shard

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	shard.ID = common.SHA256(serialized)                                        // Set shard ID
	shard.Address, err = common.SeedAddress(initializingNode.Address, shard.ID) // Generate, set address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &shard, nil // Return initialized shard
}
