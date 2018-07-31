package types

import (
	"errors"

	"github.com/mitsukomegumi/GoP2P/common"
)

// NodeDatabase - database containing list of node addresses, as well as bootstrap addresses
type NodeDatabase struct {
	Nodes *[]Node // Nodes - primary list of nodes

	AcceptableTimeout uint // AcceptableTimeout - database-wide definition for operation timeout
}

// NewDatabase - attempts creates new instance of the NodeDatabase struct
func NewDatabase(bootstrapNode *Node, timeout uint) (NodeDatabase, error) {

	db := NodeDatabase{AcceptableTimeout: timeout}

	err := db.AddNode(bootstrapNode)

	if err != nil {
		return NodeDatabase{}, err
	}

	return db, nil // No error occurred, return database
}

// AddNode - adds node to specified nodedatabase, after checking address of node
func (db *NodeDatabase) AddNode(node *Node) error {
	if !common.CheckAddress(node.Address) {
		return errors.New("invalid address")
	}

	*db.Nodes = append(*db.Nodes, *node)

	return nil
}

// RemoveNode - removes node with specified address from database
func (db *NodeDatabase) RemoveNode(address string) error {
	nodeIndex, err := db.QueryForAddress(address) // Finds index of node with address

	if err != nil { // Checks for error
		return err // Returns error
	}

	db.remove(int(nodeIndex)) // Removes value at index
	return nil                // Returns nil, no error
}

func (db *NodeDatabase) remove(s int) { // Removes address at index
	*db.Nodes = append((*db.Nodes)[:s], (*db.Nodes)[s+1:]...) // Remove index
}

// QueryForAddress - attempts to search specified node database for specified address, returns index of node
func (db *NodeDatabase) QueryForAddress(address string) (uint, error) {
	x := 0

	for x != len(*db.Nodes) { // Wait until entire db has been queried
		if address == (*db.Nodes)[x].Address { // Check for match
			return uint(x), nil // If provided value matches value of node in list, return index
		}
		x++ // Increment index
	}

	return 0, errors.New("no value found")
}
