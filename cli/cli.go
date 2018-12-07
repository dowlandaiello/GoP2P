package cli

import (
	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// NewNode - simple wrapper for initializing a new node
func NewNode() (*node.Node, error) {
	address, err := common.GetExtIPAddrWithoutUPnP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		var gErr error // Init err

		address, gErr = common.GetExtIPAddrWithUPnP() // Attempt to fetch address with UPnP

		if gErr != nil { // Check second try for errors
			return nil, gErr
		}
	}

	common.Printf("found address %s", address) // Log found address

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil { // Check for errors
		return nil, err
	}

	return &node, nil // No error occurred, return node
}

// AttachNode - attempt to attach to saved node in current working directory
func AttachNode() (*node.Node, error) {
	currentDir, err := common.GetCurrentDir() // Fetch current dir

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	readNode, err := node.ReadNodeFromMemory(currentDir) // Attempt to read serialized node

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	return readNode, nil // Return read node
}
