package node

import (
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
)

// Node - abstract struct containing metadata for a node
type Node struct {
	Address      string                   `json:"IP address"`  // Node's IP address
	Reputation   uint                     `json:"reputation"`  // Node's reputation (used for node finding algorithm)
	LastPingTime time.Time                `json:"ping"`        // Last time that the node was pinged successfully (also used for node finding algorithm)
	IsBootstrap  bool                     `json:"is boostrap"` // Value used for checking whether or not a specific node is a bootstrap node (again, used for node finding algorithm)
	Environment  *environment.Environment `json:"environment"` // Used for variable storage and referencing
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewNode - create new instance of node struct, with address specified
func NewNode(address string, isBootstrap bool) (Node, error) {
	environment, err := environment.NewEnvironment() // Create new environment

	if err != nil { // Check for errors
		return Node{}, err // Return error
	}

	if address == "" { // Check for invalid address
		return Node{}, errors.New("invalid init values") // Return error
	}

	node := Node{Address: address, Reputation: 0, IsBootstrap: isBootstrap, Environment: environment} // Creates new node instance with specified address

	err = common.CheckAddress(node.Address) // Verify address

	if err != nil { // If node address is invalid, return error
		return Node{}, err // Returns nil node, error
	}

	node.LastPingTime = common.GetCurrentTime() // Since node address is valid, add current time as last ping time
	node.Reputation += common.NodeAvailableRep

	return node, nil // No error occurred, return nil
}

// StartListener - attempt to listen on specified port, return new listener
func (node *Node) StartListener(port int) (*net.Listener, error) {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port)) // Listen on port

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &ln, nil // No error occurred, return listener
}

/*
	END EXPORTED METHODS:
*/
