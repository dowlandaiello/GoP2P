package connection

import "github.com/mitsukomegumi/GoP2P/types/node"

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

// Connection - abstract container for Golang connection type, contains metadata, routing parameteres
type Connection struct {
	DestinationNode    string     `json:"destination node"`  // Node to contact
	InitializationNode *node.Node `json:"initializing node"` // Node initializing connection

	Metadata Metadata `json:"meta"` // Connection metadata
	Data     []byte   `json:"data"` // Actual data being transmitted

	ConnectionType  string  `json:"type"` // Type of connection
	ConnectionStack []Event `json:"stack"`
}

// Event - container holding metadata concerning a direction given from a peer to another
type Event struct {
	EventType string `json:"type"`
}

// Metadata - container for connection metadata
type Metadata struct {
}
