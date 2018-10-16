package database

import (
	"errors"
	"reflect"

	"github.com/mitsukomegumi/GoP2P/types/command"

	"github.com/mitsukomegumi/GoP2P/types/connection"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// NodeDatabase - database containing list of node addresses, as well as bootstrap addresses
type NodeDatabase struct {
	Nodes *[]node.Node `json:"nodes"` // Nodes - primary list of nodes

	NetworkAlias string `json:"network"`   // NetworkAlias - network 'name', used for identifying a common protocol
	NetworkID    uint   `json:"networkID"` // NetworkID - integer used for identifying common network

	AcceptableTimeout uint `json:"db-wide timeout"` // AcceptableTimeout - database-wide definition for operation timeout
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewDatabase - attempts creates new instance of the NodeDatabase struct
func NewDatabase(bootstrapNode *node.Node, networkName string, networkID uint, acceptableTimeout uint) (NodeDatabase, error) {
	db := NodeDatabase{AcceptableTimeout: acceptableTimeout, NetworkAlias: networkName, NetworkID: networkID} // Create empty database with specified timeout

	err := db.AddNode(bootstrapNode) // Attempt to add bootstrapnode

	if err != nil { // Check for errors
		return NodeDatabase{}, err // Return empty node database, error
	}

	return db, nil // No error occurred, return database
}

/* BEGIN NODE METHODS */

// AddNode - adds node to specified nodedatabase, after checking address of node
func (db *NodeDatabase) AddNode(destNode *node.Node) error {
	err := common.CheckAddress(destNode.Address) // Attempt to check specified address

	if err != nil { // Check for invalid address
		return err // Return new error
	}

	if reflect.ValueOf(db.Nodes).IsNil() { // Check if node array is nil
		db.Nodes = &[]node.Node{*destNode} // Initialize empty array with new array composed of destNode
	} else { // Array is not nil
		*db.Nodes = append(*db.Nodes, *destNode) // Append node to node list
	}

	go db.UpdateRemoteDatabase() // Update remote database instances

	return nil // No error occurred, return nil
}

// RemoveNode - removes node with specified address from database
func (db *NodeDatabase) RemoveNode(address string) error {
	nodeIndex, err := db.QueryForAddress(address) // Finds index of node with address

	if err != nil { // Checks for error
		return err // Returns error
	}

	db.remove(int(nodeIndex)) // Removes value at index

	return nil // Returns nil, no error
}

// QueryForAddress - attempts to search specified node database for specified address, returns index of node
func (db *NodeDatabase) QueryForAddress(address string) (uint, error) {
	for x := 0; x != len(*db.Nodes); x++ { // Wait until entire db has been queried
		if address == (*db.Nodes)[x].Address { // Check for match
			return uint(x), nil // If provided value matches value of node in list, return index
		}
	}

	return 0, errors.New("no value found") // Could not find index of address, return new error
}

// UpdateRemoteDatabase - push database changes to remote network nodes
func (db *NodeDatabase) UpdateRemoteDatabase() error {
	serializedDb, err := common.SerializeToBytes(*db) // Serialize database to bytes

	if err != nil { // Check for errors
		return err // Return found error
	}

	for node := range *db.Nodes { // Iterate over nodes
		go common.SendBytes(serializedDb, (*db.Nodes)[node].Address) // Send database to node
	}

	return nil // No error occurred, return nil
}

// JoinDatabase - attempt to insert local node data into remote database instance
func JoinDatabase(bootstrapAddress string, databasePort uint, databaseAlias string) error {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return err // Return found error
	}

	localNode, err := node.ReadNodeFromMemory(currentDir) // Attempt to read node from current dir

	if err != nil { // Check for errors
		return err // Return found error
	}

	db, err := FetchRemoteDatabase(bootstrapAddress, databasePort, databaseAlias) // Fetch remote database

	if err != nil { // Check for errors
		return err // Return found error
	}

	err = db.AddNode(localNode) // Add local node

	if err != nil { // Check for errors
		return err // Return found error
	}

	err = db.UpdateRemoteDatabase() // Update remote database instances

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}

// FetchRemoteDatabase - attempt to fetch working copy of remote database
func FetchRemoteDatabase(bootstrapAddress string, databasePort uint, databaseAlias string) (*NodeDatabase, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	localNode, err := node.ReadNodeFromMemory(currentDir) // Attempt to read node from current dir

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	resolution, err := connection.NewResolution([]byte("dbFetchRequest"), "dbFetchRequest")

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	command, err := command.NewCommand("QueryType", command.NewModifierSet(databaseAlias+"NodeDatabase", nil, nil)) // Init command

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	event, err := connection.NewEvent("fetch", *resolution, command, &node.Node{Address: bootstrapAddress}, int(databasePort)) // Init event

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	connection, err := connection.NewConnection(localNode, &node.Node{Address: bootstrapAddress}, int(databasePort), []byte("dbFetchRequest"), "relay", []connection.Event{*event})

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	result, err := connection.Attempt() // Attempt connection

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	db, err := FromBytes(result) // Convert read bytes to database

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	return db, nil // No error occurred, return nil
}

/* END NODE METHODS */

/*
	END EXPORTED METHODS
*/

/*
	BEGIN INTERNAL METHODS:
*/

func (db *NodeDatabase) remove(s int) { // Removes address at index
	*db.Nodes = append((*db.Nodes)[:s], (*db.Nodes)[s+1:]...) // Remove index
}

/*
	END INTERNAL METHODS
*/
