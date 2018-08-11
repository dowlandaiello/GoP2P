package cli

import (
	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// NewNode - simple wrapper for initializing a new node
func NewNode() (*node.Node, error) {
	address, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		var gErr error // Init err

		address, gErr = common.GetExtIPAddrWithoutUpNP() // Attempt to fetch address without UpNP

		if gErr != nil { // Check second try for errors
			return nil, gErr
		}
	}

	node, err := node.NewNode(address, true) // Attempt to create new node

	if err != nil { // Check for errors
		return nil, err
	}

	currentDir, err := common.GetCurrentDir() // Fetch current directory

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	err = node.WriteToMemory(currentDir) // Attempt to write node to memory

	if err != nil { // Check for errors
		return nil, err // Return found error
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
