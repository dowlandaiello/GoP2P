package common

import (
	"path/filepath"
	"testing"
)

/*
	BEGIN EXPORTED METHODS:
*/

// TestWriteGob - test functionality of WriteGob() function
func TestWriteGob(t *testing.T) {
	dir, err := GetCurrentDir() // Get working directory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
		t.FailNow()           // Panic
	}

	err = WriteGob(dir+filepath.FromSlash("/test.gob"), "test") // Write gob to directory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
		t.FailNow()           // Panic
	}

	t.Logf("gob written '%s'", "test")
}

// TestReadGob - test functionality of ReadGob() function
func TestReadGob(t *testing.T) {
	dir, err := GetCurrentDir() // Get working directory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
		t.FailNow()           // Panic
	}

	err = WriteGob(dir+filepath.FromSlash("/test.gob"), "test") // Write gob to directory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
		t.FailNow()           // Panic
	}

	var input string // Create buffer for read gob

	err = ReadGob(dir+filepath.FromSlash("/test.gob"), &input) // Attempt to read gob at directory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
		t.FailNow()           // Panic
	}

	t.Logf("gob read '%s'", input) // Log read gob
}

// TestSerializeToBytes - test functionality of SerializeToBytes() function
func TestSerializeToBytes(t *testing.T) {
	obj := "test" // Create temporary testing object

	byteVal, err := SerializeToBytes(obj) // Attempt to serialize

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found serialized value %s", byteVal) // Log success
}

func TestSerializeToString(t *testing.T) {
	obj := []string{"test", "test", "test"} // Create temporary testing object

	stringVal, err := SerializeToString(obj) // Attempt to serialize

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found serialized value %s", stringVal) // Log success
}

/*
	END EXPORTED METHODS
*/
