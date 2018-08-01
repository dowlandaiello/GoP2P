package connection

import (
	"errors"
	"reflect"
	"strings"

	"github.com/mitsukomegumi/GoP2P/common"
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

	// AvailableEventTypes - global preset list of directions that an individual node can give to another
	AvailableEventTypes = []string{"push", "fetch"}

	/*
		Push - request for a node to push information to a certain location
		Fetch - request for a node to fetch information from a certain location
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

// Event - container holding metadata concerning a direction given from a peer to another
type Event struct {
	EventType string `json:"type"`

	Data []byte `json:"data"` // Data being transmitted

	SourceNode      *node.Node `json:"source"`      // Initialization node
	DestinationNode *node.Node `json:"destination"` // Node to contact
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

// NewEvent - creates new Event{} instance with specified data, peers
func NewEvent(eventType string, data []byte, sourceNode *node.Node, destinationNode *node.Node) (*Event, error) {
	if strings.ToLower(eventType) != "push" && strings.ToLower(eventType) != "fetch" { // Check for invalid types
		return &Event{}, errors.New("invalid event") // Error occurred, return nil, error
	} else if reflect.ValueOf(sourceNode).IsNil() || reflect.ValueOf(destinationNode).IsNil() { // Check for invalid peer values
		return &Event{}, errors.New("invalid peer value") // Error occurred, return nil, error
	}

	return &Event{EventType: eventType, Data: data, SourceNode: sourceNode, DestinationNode: destinationNode}, nil // Return initialized event
}

// Attempt - attempts to carry out connection, if connection stack is provided, begins to iterate through list
func (connection *Connection) Attempt() {

}

// Attempt - attempts to carry out event
func (event *Event) Attempt() error {
	isSelf, err := event.checkIsSelf()

	if err != nil { // Check for errors
		return err // Return error
	}

	if isSelf == true { // Handle different types of event
		event.attemptIsSelf()
	} else {
		event.attemptExternal()
	}

	return nil
}

func (event *Event) attemptIsSelf() error {
	return nil
}

func (event *Event) attemptExternal() error {
	return nil
}

// CheckIsSelf - check whether or not an event is being passed to a node with the intention of receiving data from that node
func (event *Event) checkIsSelf() (bool, error) {
	selfAddr, err := common.GetExtIPAddrWithUpNP() // Attempt to fetch current external IP address

	if err != nil { // Check for errors
		selfAddr, err = common.GetExtIPAddrWithoutUpNP() // Fetch IP in a safer manner

		if err != nil { // Check for errors
			return false, err // Return error
		}
	}

	return event.SourceNode.Address == selfAddr, nil // No error occurred, return nil error
}

/*
	END EXPORTED METHODS
*/
