package shard

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// Shard - container holding shard metadata
type Shard struct {
	Nodes      *[]node.Node `json:"nodes"`       // Nodes - primary list of nodes
	ChildNodes *[]node.Node `json:"allChildren"` // ChildNodes - list of all child nodes (recursively includes nodes in child shards, not just direct children)

	ShardRoot   *Shard `json:"-"`      // ShardRoot - root shard of shard tree
	Root        bool   `json:"isRoot"` // Root - is root
	ParentShard *Shard `json:"-"`      // ParentShard - parent of shard

	Siblings *[]*Shard `json:"siblings"` // Siblings - shard-level siblings

	ChildShards []*Shard `json:"child shards"` // ChildShards - shards created as children of shard

	Age    uint64    `json:"age"`           // Age - shard age
	Origin time.Time `json:"creation time"` // Origin - time shard created

	Address string `json:"address"` // Address - addressable internet protocol ID used for shard-level communications

	ID string `json:"id"` // ID - hash of Shard contents
}

// NewShard - initialize new shard
func NewShard(initializingNode *node.Node) (*Shard, error) {
	shard := Shard{Nodes: &[]node.Node{*initializingNode}, ChildNodes: &[]node.Node{*initializingNode}, ChildShards: []*Shard{}, Origin: time.Now().UTC(), Address: (*initializingNode).Address} // Initialize shard

	serialized, err := common.SerializeToBytes(shard) // Serialize shard

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	shard.ID = common.Sha3(serialized)                                             // Set shard ID
	shard.Address, err = common.SeedAddress((*initializingNode).Address, shard.ID) // Generate, set address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &shard, nil // Return initialized shard
}

// NewShardWithNodes - initialize new shard with child nodes
func NewShardWithNodes(initializingNodes *[]node.Node) (*Shard, error) {
	shard := Shard{Nodes: initializingNodes, ChildNodes: initializingNodes, ChildShards: []*Shard{}, Origin: time.Now().UTC(), Address: ""} // Initialize shard

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

// Shard - exponentially shard specified shard into child shards
func (shard *Shard) Shard(exponent uint) error {
	totalShards := math.Pow(float64(exponent), float64(exponent)) // Calculate total shards

	if totalShards > float64(len(*shard.ChildNodes)) { // Check for invalid node count
		return errors.New("shard smaller than exponential shard count") // Return found error
	}

	if reflect.ValueOf(shard.ParentShard).IsNil() { // Check is root
		(*shard).Root = true // Set root
	}

	lastShard := shard // Set last shard

	for x := 0; x != int(exponent); x++ {
		for z := 0; z != int(exponent); z++ {
			foundNodes := (*shard.ChildNodes)[(z * x):((z * x) + int(exponent))] // Fetch nodes in shard

			newShard, err := NewShardWithNodes(&foundNodes) // Init shard

			if err != nil { // Check for errors
				return err // Return found error
			}

			(*newShard).ParentShard = lastShard                                // Set parent
			(*newShard).ShardRoot = shard                                      // Set shard root
			(*lastShard).ChildShards = append(lastShard.ChildShards, newShard) // Append initialized shard

			lastShard = newShard // Set last shard
		}
	}

	return nil // No error occurred, return nil
}

// LogShard - serialize and print contents of entire shard
func (shard *Shard) LogShard() error {
	marshaledVal, err := json.MarshalIndent(*shard, "", "  ") // Marshal shard

	if err != nil { // Check for errors
		return err // Return found error
	}

	fmt.Println("\n" + string(marshaledVal)) // Log marshaled val

	return nil // No error occurred, return nil
}

// CalculateQuadraticExponent - returns exponential value of exponent to exponent exponent times
func CalculateQuadraticExponent(exponent int) float64 {
	val := float64(exponent) // Init buffer

	for x := 0; x != exponent; x++ {
		val = math.Pow(float64(val), float64(exponent)) // Set to exponent
	}

	return val // Return calculated val
}
