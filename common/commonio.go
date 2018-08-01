package common

import (
	"encoding/gob"
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

/*
	END EXPORTED METHODS
*/
