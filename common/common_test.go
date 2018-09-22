package common

import (
	"strings"
	"testing"
)

// TestParseStringParams - test functionality of ParseStringParams() function
func TestParseStringParams(t *testing.T) {
	input := "node.NewNode(localhost, 3000)" // Init input

	params, err := ParseStringParams(input) // Parse string params

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
		t.FailNow()           // Panic
	}

	t.Logf("found parsed params %s, %s", params[0], params[1]) // Log success
}

// TestStringStripParentheses - test functionality of StringStripParentheses() function
func TestStringStripParentheses(t *testing.T) {
	input := "node.NewNode(localhost, 3000)" // Init input

	stripped := StringStripParentheses(input) // Strip parentheses

	t.Logf("found value %s", stripped) // Log success
}

// TestCheckAddress - test functionality of CheckAddress() function
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

// TestSHA256 - test functionality of SHA256() function
func TestSHA256(t *testing.T) {
	shaVal := SHA256([]byte("test")) // Create hashed value

	if string(shaVal) == "test" { // Check for errors
		t.Errorf("SHA failed with value %s", shaVal) // Log error
		t.FailNow()                                  // Panic
	}

	t.Logf("hashed value %s", shaVal) // Log success
}
