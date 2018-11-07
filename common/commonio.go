package common

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

/*
	BEGIN EXPORTED METHODS:
*/

// WriteGob - create gob from specified object, at filePath
func WriteGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath) // Attempt to create file at path

	if err != nil { // Check for errors
		return err // Return found error
	}

	encoder := gob.NewEncoder(file) // Write to file
	err = encoder.Encode(object)    // Encode object

	if err != nil { // Check for errors
		return err // Return found error
	}

	file.Close() // Close file operation
	return err   // Return error (might be nil)
}

// ReadGob - read gob specified at path
func ReadGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath) // Attempt to open file at path

	if err != nil { // Check for errors
		return err // Return found error
	}

	decoder := gob.NewDecoder(file) // Attempt to decode gob
	err = decoder.Decode(object)    // Assign to error

	if err != nil { // Check for errors
		return err // Return found error
	}

	file.Close() // Close file
	return err   // Return error
}

// SerializeToBytes - attempt to convert specified interface to byte array
func SerializeToBytes(object interface{}) ([]byte, error) {
	serializedBuffer := new(bytes.Buffer) // Create buffer to store encoded object

	err := json.NewEncoder(serializedBuffer).Encode(object) // Attempt to encode

	if err != nil { // Check for errors
		return nil, err // Return error
	}

	return serializedBuffer.Bytes(), nil // Return serialized object
}

// InterfaceFromBytes - attempt to decode specified byte array to interface
func InterfaceFromBytes(data []byte, buffer interface{}) (interface{}, error) {
	err := json.NewDecoder(bytes.NewReader(data)).Decode(&buffer) // Attempt to read

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return buffer, nil // No error occurred, return read value
}

// SerializeToString - attempt to get string representation of specified object
func SerializeToString(object interface{}) (string, error) {
	out, err := json.Marshal(object) // Attempt to marshal object

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return string(out), nil // Return serialized value
}

// MarshalInterfaceToMap - attempt to map interface to string map
func MarshalInterfaceToMap(object interface{}) (map[string]string, error) {
	mapped := structs.Map(object) // Attempt to create interface map

	stringMap := make(map[string]string) // Init buffer

	for key, value := range mapped { // Iterate through field values
		strKey := fmt.Sprintf("%v", key)     // Gen key
		strValue := fmt.Sprintf("%v", value) // Gen value

		stringMap[strKey] = strValue // Create value at key
	}

	return stringMap, nil // Return decoded result
}

// UnmarshalInterfaceFromMap - attempt to convert specified map to interface
func UnmarshalInterfaceFromMap(object map[string]interface{}) (interface{}, error) {
	var buf interface{} // Init buffer

	err := mapstructure.Decode(object, &buf) // Attempt to decode

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return buf, nil // Return decoded result
}

// UnmarshalInterfaceFromStringMap - attempt to convert specified map to interface
func UnmarshalInterfaceFromStringMap(buffer interface{}, object map[string]string) (interface{}, error) {
	err := mapstructure.Decode(object, &buffer) // Attempt to decode

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return buffer, nil // Return decoded result
}

/*
	END EXPORTED METHODS
*/
