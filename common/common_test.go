package common

import (
	"path/filepath"
	"strings"
	"testing"
)

// TestCheckAddress - test functionality of CheckAddress() function.
func TestCheckAddress(t *testing.T) {
	err := CheckAddress("72.21.215.90") // Attempt to check the address of S3

	if err != nil && !strings.Contains(err.Error(), "root") { // Check for errors
		t.Errorf(err.Error()) // Log errors
		t.FailNow()           // Panic
	}
}

// TestGetExtIPAddrWithUpNP - test functionality of GetExtIPAddrWithUpNP() function
func TestGetExtIPAddrWithUpNP(t *testing.T) {
	ip, err := GetExtIPAddrWithUpNP() // Attempt to fetch external IP address

	if err != nil && !strings.Contains(err.Error(), "gateway found") { // Check for errors
		t.Errorf(err.Error()) // Log errors to console
		t.FailNow()           // Panic on found error
	}

	t.Logf("found address %s", ip) // Log found address
}

// TestGetExtIPAddrWithoutUpNP - test functionality of GetExtIPAddrWithoutUpNP() function
func TestGetExtIPAddrWithoutUpNP(t *testing.T) {
	ip, err := GetExtIPAddrWithoutUpNP() // Attempt to fetch external IP address

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log errors to console
		t.FailNow()           // Panic on found error
	}

	t.Logf("found address %s", ip) // Log found address
}

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
