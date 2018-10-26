package connection

import (
	"bytes"
	"encoding/json"
)

// Response - abstract container holding array of byte arrays
type Response struct {
	Val [][]byte `json:"value"`
}

// ResponseFromBytes - attempt to convert specified byte array to connection
func ResponseFromBytes(b []byte) (*Response, error) {
	b = bytes.Trim(b, "\x00") // Trim null character

	object := Response{} // Create empty instance

	// PANICS HERE: \x00 character in string literal https://forum.golangbridge.org/t/how-to-convert-utf-8-string-to-utf-16-be-string/7072

	err := json.NewDecoder(bytes.NewReader(b)).Decode(&object) // Attempt to read

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &object, nil // No error occurred, return read value
}
