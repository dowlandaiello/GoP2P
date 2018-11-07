package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/mitsukomegumi/GoP2P/common"
)

var (
	// ValidMessageTypes - definitions for valid messsage types
	ValidMessageTypes = []string{"hardfork", "update", "upgrade", "notice"}
)

// Message - string alias for database messages
type Message struct {
	Message  string `json:"message"`     // Message value
	Priority uint   `json:"priority"`    // Message priority
	Type     string `json:"messagetype"` // Message type
}

// NewMessage - attempt to initialize new message with given parameters
func NewMessage(message string, priority uint, messageType string) (*Message, error) {
	if !common.StringInSlice(ValidMessageTypes, messageType) { // Check for invalid message type
		return &Message{}, fmt.Errorf("invalid message type %s", messageType) // Return found error
	}

	return &Message{Message: message, Priority: priority, Type: messageType}, nil // Return initialized message
}

// ToBytes - attempt to serialize given message to bytes
func (message *Message) ToBytes() ([]byte, error) {
	if reflect.ValueOf(message).IsNil() { // Check for nil message
		return []byte{}, errors.New("nil input") // Return found error
	}

	return common.SerializeToBytes(*message) // Serialize message
}

// MessageFromBytes - attempt to decode message from given bytes
func MessageFromBytes(b []byte) (*Message, error) {
	if len(b) == 0 { // Check for invalid input
		return nil, errors.New("nil input") // Return found error
	}

	object := Message{} // Create empty instance

	err := json.NewDecoder(bytes.NewReader(b)).Decode(&object) // Attempt to read

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &object, nil // No error occurred, return read value
}
