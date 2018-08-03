package connection

import (
	"errors"
	"reflect"
	"strings"

	"github.com/mitsukomegumi/GoP2P/types/node"
)

var (
	// AvailableConnectionTypes - global preset list of types connections that can be made
	AvailableConnectionTypes = []string{"relay", "pointer"}

	/*
		Relay - sending information from one peer to another
		Pointer - sending metadata from one peer to another requesting for the peer to fetch information from another peer
			Example: peer one contacts peer two and asks for a block, peer two points peer one to peer three
	*/
)

// Connection - abstract container for Golang connection type, contains metadata, routing parameters
type Connection struct {
	DestinationNode    *node.Node `json:"destination node"`  // Node to contact
	InitializationNode *node.Node `json:"initializing node"` // Node initializing connection

	Data []byte `json:"data"` // Actual data being transmitted

	ConnectionType  string  `json:"type"` // Type of connection
	ConnectionStack []Event `json:"stack"`
}

// Resolution - abstract type defining how to handle and deal with a connection or event's data
type Resolution struct {
	ResolutionData []byte `json:"data"` // ResolutionData - data being passed via resolution (typically a struct)

	GuidingType interface{} `json:"guide"` // GuidingType - guiding struct to map resolution fields
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewConnection - creates new Connection{} instance with specified data, peers
func NewConnection(sourceNode *node.Node, destinationNode *node.Node, data []byte, connectionType string, connectionStack []Event) (*Connection, error) {
	if strings.ToLower(connectionType) != "relay" && strings.ToLower(connectionType) != "pointer" { // Check connection type is valid
		return &Connection{}, errors.New("invalid connection type") // Error occurred, return nil
	} else if reflect.ValueOf(destinationNode).IsNil() || reflect.ValueOf(sourceNode).IsNil() { // Check that peer values aren't nil
		return &Connection{}, errors.New("invalid peer value") // Peer values nil, return nil constructor
	} else if len(data) == 0 { // Check that data is being passed trough
		return &Connection{}, errors.New("invalid data") // Return error
	}

	return &Connection{DestinationNode: destinationNode, InitializationNode: sourceNode, Data: data, ConnectionType: connectionType, ConnectionStack: connectionStack}, nil // No error occurred, return correctly initialized Connection
}

// NewResolution - attempt to create new instance of the Resolution struct with specified initializers
func NewResolution(data []byte, guidingType interface{}) (*Resolution, error) {
	if len(data) == 0 { // Check for invalid data
		return &Resolution{}, errors.New("nil value found") // Return found error
	}

	return &Resolution{ResolutionData: data, GuidingType: guidingType}, nil // No error occurred, return initialized Resolution
}

// Attempt - attempts to carry out connection, if event stack is provided, begins to iterate through list
func (connection *Connection) Attempt() {

}

/*
	END EXPORTED METHODS
*/
