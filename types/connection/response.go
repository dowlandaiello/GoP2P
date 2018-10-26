package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Response - abstract container holding array of byte arrays
type Response struct {
	Val [][]byte `json:"value"`
}

// ResponseFromBytes - attempt to convert specified byte array to connection
func ResponseFromBytes(b []byte) (*Response, error) {
	object := Response{} // Create empty instance

	fmt.Println("test")

	err := json.NewDecoder(bytes.NewReader(b)).Decode(&object) // Attempt to read

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &object, nil // No error occurred, return read value
}
