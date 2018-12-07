package connection

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

var (
	// AvailableConnectionTypes - global preset list of types connections that can be made
	AvailableConnectionTypes = []string{"relay"}

	/*
		Relay - sending information from one peer to another
	*/
)

// Connection - abstract container for Golang connection type, contains metadata, routing parameters
type Connection struct {
	DestinationNode    *node.Node `json:"destination node"`  // Node to contact
	InitializationNode *node.Node `json:"initializing node"` // Node initializing connection

	Data []byte `json:"data"` // Actual data being transmitted

	Port int `json:"port"`

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
func NewConnection(sourceNode *node.Node, destinationNode *node.Node, port int, data []byte, connectionType string, connectionStack []Event) (*Connection, error) {
	if strings.ToLower(connectionType) != "relay" && strings.ToLower(connectionType) != "pointer" { // Check connection type is valid
		return &Connection{}, errors.New("invalid connection type") // Error occurred, return nil
	} else if reflect.ValueOf(destinationNode).IsNil() || reflect.ValueOf(sourceNode).IsNil() { // Check that peer values aren't nil
		return &Connection{}, errors.New("invalid peer value") // Peer values nil, return nil constructor
	} else if len(data) == 0 { // Check that data is being passed trough
		return &Connection{}, errors.New("invalid data") // Return error
	}

	return &Connection{DestinationNode: destinationNode, Port: port, InitializationNode: sourceNode, Data: data, ConnectionType: connectionType, ConnectionStack: connectionStack}, nil // No error occurred, return correctly initialized Connection
}

// NewResolution - attempt to create new instance of the Resolution struct with specified initializers
func NewResolution(data []byte, guidingType interface{}) (*Resolution, error) {
	if len(data) == 0 { // Check for invalid data
		return &Resolution{}, errors.New("nil value found") // Return found error
	}

	return &Resolution{ResolutionData: data, GuidingType: guidingType}, nil // No error occurred, return initialized Resolution
}

// Attempt - attempts to carry out connection, if event stack is provided, begins to iterate through list
func (connection *Connection) Attempt() ([]byte, error) {
	return connection.attempt() // Found connection stack, handle respectively
}

// AttemptVariable - attempts to carry out connection, returning variable response
func (connection *Connection) AttemptVariable() (*environment.Variable, error) {
	response, err := connection.attempt() // Attempt connection

	if err != nil { // Check for errors
		return &environment.Variable{}, err // Return found error
	}

	decodedResponse, err := ResponseFromBytes(response) // Fetch decoded result

	if err != nil { // Check for errors
		return &environment.Variable{}, err // Return found error
	}

	return environment.VariableFromBytes(decodedResponse.Val[0]) // Return final decoded response
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// attempt - attempt connection
func (connection *Connection) attempt() ([]byte, error) {
	common.Println("-- CONNECTION -- attempting connection to peer with address " + connection.DestinationNode.Address) // Log connection

	serializedConnection, err := common.SerializeToBytes(*connection) // Serialize connection

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	result, err := common.SendBytesResult(serializedConnection, connection.DestinationNode.Address+":"+strconv.Itoa(connection.Port)) // Attempt to send event

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	/*
		TODO: fix nil read data
	*/

	return result, nil // No error occurred, return nil
}

/* END INTERNAL METHODS */
