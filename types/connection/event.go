package connection

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/command"
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

	Command *command.Command `json:"command"` // Action for destination node to carry out

	DestinationNode *node.Node `json:"destination"` // Node to contact

	Port int `json:"port"`
}

/*
	BEGIN EXPORTED METHODS:
*/

// NewEvent - creates new Event{} instance with specified resolution, peers, command
func NewEvent(eventType string, resolution Resolution, command *command.Command, destinationNode *node.Node, port int) (*Event, error) {
	if strings.ToLower(eventType) != "push" && strings.ToLower(eventType) != "fetch" { // Check for invalid types
		return &Event{}, errors.New("invalid event type") // Error occurred, return nil, error
	} else if reflect.ValueOf(destinationNode).IsNil() { // Check for invalid peer values
		return &Event{}, errors.New("invalid peer value") // Error occurred, return nil, error
	} else if reflect.ValueOf(command).IsNil() { // Check for nil command
		return &Event{}, errors.New("invalid command") // Error occurred, return nil, error
	}

	return &Event{EventType: eventType, Resolution: resolution, Command: command, DestinationNode: destinationNode, Port: port}, nil // Return initialized event
}

// Attempt - attempts to carry out event
func (event *Event) Attempt() ([]byte, error) {
	return event.attempt() // attempt
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// attempt - wrapper
func (event *Event) attempt() ([]byte, error) {
	serializedEvent, err := common.SerializeToBytes(event) // Serialize event

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	result, err := common.SendBytesResult(serializedEvent, event.DestinationNode.Address+":"+strconv.Itoa(event.Port)) // Attempt to send event

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return result, nil // No error occurred, return nil
}

/* END INTERNAL METHODS */
