package connection

import (
	"bytes"
	"encoding/json"
)

/*
	BEGIN EXPORTED METHODS:
*/

// FromBytes - attempt to convert specified byte array to connection
func FromBytes(b []byte) (*Connection, error) {
	object := Connection{} // Create empty instance

	err := json.NewDecoder(bytes.NewReader(b)).Decode(&object) // Attempt to read

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &object, nil // No error occurred, return read value
}

/*
	END EXPORTED METHODS
*/
