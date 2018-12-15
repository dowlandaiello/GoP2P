package connection

import (
	"bytes"
	"encoding/json"
	"unicode/utf8"

	"golang.org/x/text/encoding/unicode"
)

// Response - abstract container holding array of byte arrays
type Response struct {
	Val [][]byte `json:"value"`
}

// ResponseFromBytes - attempt to convert specified byte array to connection
func ResponseFromBytes(b []byte) (*Response, error) {
	var err error // Init error

	if !utf8.Valid(b) { // Check for non-valid in UTF-8
		b, err = unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder().Bytes(b) // Convert to utf-8

		if err != nil { // Check for errors
			return nil, err // Return found error
		}
	} else if bytes.Contains(b, []byte("\x00")) { // Check for \x00 null chars
		b = bytes.Trim(b, "\x00") // Trim \x00 null char
	}

	object := Response{} // Create empty instance

	err = json.NewDecoder(bytes.NewReader(b)).Decode(&object) // Attempt to read

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &object, nil // No error occurred, return read value
}
