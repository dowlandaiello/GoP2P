package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/command"
	"github.com/mitsukomegumi/GoP2P/types/connection"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
	"github.com/mitsukomegumi/GoP2P/types/shard"
)

// NodeDatabase - database containing list of node addresses, as well as bootstrap addresses
type NodeDatabase struct {
	Nodes *[]node.Node `json:"nodes"` // Nodes - primary list of nodes

	Shards *[]shard.Shard `json:"shards"` // Shards - primary list of shards

	NetworkAlias string `json:"network"`   // NetworkAlias - network 'name', used for identifying a common protocol
	NetworkID    uint   `json:"networkID"` // NetworkID - integer used for identifying common network

	HashedNetworkMessageKey string // HashedNetworkMessageKey - key used for network-wide messages

	AcceptableTimeout uint `json:"db-wide timeout"` // AcceptableTimeout - database-wide definition for operation timeout
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewDatabase - attempts creates new instance of the NodeDatabase struct
func NewDatabase(bootstrapNode *node.Node, networkName string, networkID uint, acceptableTimeout uint, privateNetworkKey string) (NodeDatabase, error) {
	db := NodeDatabase{AcceptableTimeout: acceptableTimeout, NetworkAlias: networkName, NetworkID: networkID, HashedNetworkMessageKey: common.Sha3([]byte(privateNetworkKey + networkName))} // Create empty database with specified timeout

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

/* END NODE METHODS */

/* BEGIN SHARD METHODS */

// AddShard - attempt to append shard to current NodeDatabase
func (db *NodeDatabase) AddShard(destinationShard *shard.Shard) error {
	if reflect.ValueOf(destinationShard).IsNil() || len(*destinationShard.ChildNodes) == 0 || destinationShard.Address == "" { // Check for invalid shard
		return errors.New("invalid shard") // Return found error
	}

	for _, node := range *destinationShard.Nodes { // Iterate through nodes in database
		_, err := db.QueryForAddress(node.Address) // Check if node exists in database

		if err != nil { // Check for errors while querying for address
			go db.AddNode(&node) // Add node
		}
	}

	if db.Shards != nil { // Check for non-nil shards
		for _, xShard := range *db.Shards { // Iterate through shards in database
			_, err := db.QueryForShardAddress(destinationShard.Address) // Check shard in database

			if err == nil { // Check for errors
				db.RemoveShard(xShard.Address) // Remove shard
			}
		}

		*db.Shards = append(*db.Shards, *destinationShard) // Append shard
	} else {
		db.Shards = &[]shard.Shard{*destinationShard} // Initialize w/shard
	}

	err := db.UpdateRemoteDatabase() // Update remote database instances

	if err != nil { // Check for errors
		return err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Get working directory

	if err != nil { // Check for errors
		return err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Read node from working dir

	if err != nil { // Check for errors
		return err // Return found error
	}

	err = db.WriteToMemory(node.Environment) // Write to local environment

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}

// RemoveShard - removes shard with specified address from database
func (db *NodeDatabase) RemoveShard(address string) error {
	shardIndex, err := db.QueryForShardAddress(address) // Finds index of node with address

	if err != nil { // Checks for error
		return err // Returns error
	}

	db.removeShard(int(shardIndex)) // Removes value at index

	return nil // Returns nil, no error
}

/* END SHARD METHODS */

// QueryForAddress - attempts to search specified node database for specified address, returning index of node
func (db *NodeDatabase) QueryForAddress(address string) (uint, error) {
	for x := 0; x != len(*db.Nodes); x++ { // Wait until entire db has been queried
		if address == (*db.Nodes)[x].Address { // Check for match
			return uint(x), nil // If provided value matches value of node in list, return index
		}
	}

	return 0, errors.New("no value found") // Could not find index of address, return new error
}

// QueryForShardAddress - attempts to search specified node database for specified address, returning index of shard
func (db *NodeDatabase) QueryForShardAddress(address string) (uint, error) {
	if db.Shards != nil {
		for x := 0; x != len(*db.Shards); x++ { // Wait until entire db has been queried
			if address == (*db.Shards)[x].Address { // Check for match
				return uint(x), nil // Return matching index
			}
		}

		return 0, errors.New("no value found") // Could not find index of address, return new error
	}

	return 0, fmt.Errorf("no shards in db %v", db) // Return no shards error
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

	err = db.WriteToMemory((*localNode).Environment) // Write db to memory

	if err != nil { // Check for errors
		return err // Return found error
	}

	err = localNode.WriteToMemory(currentDir) // Write node to memory

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

	resolution, err := connection.NewResolution([]byte("dbFetchRequest"), "dbFetchRequest") // Init resolution

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

	conn, err := connection.NewConnection(localNode, &node.Node{Address: bootstrapAddress}, int(databasePort), []byte("dbFetchRequest"), "relay", []connection.Event{*event}) // Init connection

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	resultBytes, err := conn.Attempt() // Attempt connection

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	decodedResponse, err := connection.ResponseFromBytes(resultBytes) // Fetch decoded result

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	decodedVariable, err := environment.VariableFromBytes(decodedResponse.Val[0]) // Attempt to decode response

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	db, err := FromBytes(decodedVariable.VariableData) // Convert read variable to database

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	err = db.WriteToMemory(localNode.Environment) // Write db to memory

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	err = localNode.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	return db, nil // No error occurred, return nil
}

// SendDatabaseMessage - send announcement message to all nodes in network
func (db *NodeDatabase) SendDatabaseMessage(message *Message, messageKey string, databasePort uint) error {
	if common.Sha3([]byte(messageKey+db.NetworkAlias)) != db.HashedNetworkMessageKey { // Check for matching message private key
		return errors.New("invalid message private key") // Return found error
	}

	byteVal, err := message.ToBytes() // Serialize to bytes

	if err != nil { // Check for errors
		return err // Return found error
	}

	finished := make(chan bool) // Init finished buffer

	for _, node := range *db.Nodes { // Iterate through nodes
		go common.SendBytesAsyncRoutine(byteVal, node.Address+":"+strconv.Itoa(int(databasePort)), finished) // Send message
	}

	<-finished // Check finished

	return nil // No error occurred, return nil
}

// LogDatabase - serialize and print contents of entire database
func (db *NodeDatabase) LogDatabase() error {
	marshaledVal, err := json.MarshalIndent(*db, "", "  ") // Marshal database

	if err != nil { // Check for errors
		return err // Return found error
	}

	fmt.Println("\n" + string(marshaledVal)) // Log marshaled val

	return nil // No error occurred, return nil
}

/*
	END EXPORTED METHODS
*/

/*
	BEGIN INTERNAL METHODS:
*/

func (db *NodeDatabase) remove(s int) { // Removes address at index
	*db.Nodes = append((*db.Nodes)[:s], (*db.Nodes)[s+1:]...) // Remove index
}

func (db *NodeDatabase) removeShard(s int) { // Removes address at index
	*db.Shards = append((*db.Shards)[:s], (*db.Shards)[s+1:]...) // Remove index
}

/*
	END INTERNAL METHODS
*/
