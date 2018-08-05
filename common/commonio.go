package common

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"os"
)

/*
	BEGIN EXPORTED METHODS:
*/

// WriteGob - create gob from specified object, at filePath
func WriteGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath) // Attempt to create file at path

	if err == nil { // Check for nil error
		encoder := gob.NewEncoder(file) // Write to file
		encoder.Encode(object)          // Encode object
	}

	file.Close() // Close file operation
	return err   // Return error (might be nil)
}

// ReadGob - read gob specified at path
func ReadGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath) // Attempt to open file at path

	if err == nil { // Check for nil error
		decoder := gob.NewDecoder(file) // Attempt to decode gob
		err = decoder.Decode(object)    // Assign to error
	}

	file.Close() // Close file
	return err   // Return error
}

// SerializeToBytes - attempt to convert specified interface to byte array
func SerializeToBytes(object interface{}) ([]byte, error) {
	serializedBuffer := new(bytes.Buffer)                   // Create buffer to store encoded object
	err := json.NewEncoder(serializedBuffer).Encode(object) // Attempt to encode

	if err != nil { // Check for errors
		return nil, err // Return error
	}

	return serializedBuffer.Bytes(), nil // Return serialized object
}

// MarshalBytes - attempt to convert specified byte array to interface
func MarshalBytes(b []byte, object interface{}) (*interface{}, error) {
	buf := bytes.NewBuffer(b)

	err := binary.Read(buf, binary.BigEndian, &object)

	if err != nil {
		return nil, err
	}

	return &object, nil
}

// SerializeToString - attempt to get string representation of specified object
func SerializeToString(object interface{}) (string, error) {
	out, err := json.Marshal(object) // Attempt to marshal object

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return string(out), nil // Return serialized value
}

/*
	END EXPORTED METHODS
*/
