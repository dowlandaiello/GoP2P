package node

import (
	"time"

	"github.com/mitsukomegumi/GoP2P/common"
)

// Node - abstract struct containing metadata for a node
type Node struct {
	Address      string    `json:"IP address"`  // Node's IP address
	Reputation   uint      `json:"reputation"`  // Node's reputation (used for node finding algorithm)
	LastPingTime time.Time `json:"ping"`        // Last time that the node was pinged successfully (also used for node finding algorithm)
	IsBootstrap  bool      `json:"is boostrap"` // Value used for checking whether or not a specific node is a bootstrap node (again, used for node finding algorithm)
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewNode - create new instance of node struct, with address specified
func NewNode(address string, isBootstrap bool) (Node, error) {
	node := Node{Address: address, Reputation: 0, IsBootstrap: isBootstrap} // Creates new node instance with specified address

	err := common.CheckAddress(address)

	if err != nil { // If node address is invalid, return error
		return Node{}, err // Returns nil node, error
	}

	node.LastPingTime = common.GetCurrentTime() // Since node address is valid, add current time as last ping time
	node.Reputation += common.NodeAvailableRep

	return node, nil // No error occurred, return nil
}

/*
	END EXPORTED METHODS:
*/
