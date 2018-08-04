package connection

import (
	"errors"
	"reflect"
	"strings"

	"github.com/mitsukomegumi/GoP2P/types/node"
)

var (
	// AvailableEventTypes - global preset list of directions that an individual node can give to another
	AvailableEventTypes = []string{"push", "fetch"}

	/*
		Push - request for a node to push information to a certain location
		Fetch - request for a node to fetch information from a certain location
	*/
)

// Event - container holding metadata concerning a direction given from a peer to another
type Event struct {
	EventType string `json:"type"`

	Resolution Resolution `json:"resolution"` // Data being transmitted

	Command string `json:"command"` // Action for destination node to carry out

	DestinationNode *node.Node `json:"destination"` // Node to contact
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewEvent - creates new Event{} instance with specified resolution, peers
func NewEvent(eventType string, resolution Resolution, command string, destinationNode *node.Node) (*Event, error) {
	if strings.ToLower(eventType) != "push" && strings.ToLower(eventType) != "fetch" { // Check for invalid types
		return &Event{}, errors.New("invalid event") // Error occurred, return nil, error
	} else if reflect.ValueOf(destinationNode).IsNil() { // Check for invalid peer values
		return &Event{}, errors.New("invalid peer value") // Error occurred, return nil, error
	}

	return &Event{EventType: eventType, Resolution: resolution, Command: command, DestinationNode: destinationNode}, nil // Return initialized event
}

// Attempt - attempts to carry out event
func (event *Event) Attempt() error {
	err := event.attempt() // attempt

	if err != nil { // Check for errors
		return err // Return error
	}

	return nil // No error occurred, return nil
}

// attempt - wrapper
func (event *Event) attempt() error {
	return nil // No error occurred, return nil
}

/*
	END EXPORTED METHODS
*/
