package common

import (
	"strings"
	"testing"
)

// TestSeedAddress - test functionality of SeedAddress() function
func TestSeedAddress(t *testing.T) {
	seed := "despacito"
	shardID := Keccak256([]byte("despacito"))

	address, err := SeedAddress(seed, shardID) // Seed address

	if err != nil {
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("seeded address %s", address) // Log success
}

// TestParseStringMethodCall - test functionality of ParseStringMethodCall() function
func TestParseStringMethodCall(t *testing.T) {
	input := "node.NewNode(localhost, 3000)" // Init input

	receiver, methodName, params, err := ParseStringMethodCall(input) // Parse string method call

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log error
		t.FailNow()           // Panic
	}

	t.Logf("found parsed method call %s, %s, %s, %s", receiver, methodName, params[0], params[1]) // Log success
}

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

// TestStringStripReceiverCall - test functionality of StripReceiverCall() function
func TestStringStripReceiverCall(t *testing.T) {
	input := "node.NewNode(localhost, 3000)" // Init input

	stripped := StringStripReceiverCall(input) // Parse string params

	t.Logf("found stripped %s", stripped) // Log success
}

// TestStringStripParentheses - test functionality of StringStripParentheses() function
func TestStringStripParentheses(t *testing.T) {
	input := "node.NewNode(localhost, 3000)" // Init input

	stripped := StringStripParentheses(input) // Strip parentheses

	t.Logf("found value %s", stripped) // Log success
}

// TestStringFetchCallReceiver - test functionality of StringFetchCallReceiver() method
func TestStringFetchCallReceiver(t *testing.T) {
	input := "node.NewNode(localhost, 3000)" // Init input

	receiver := StringFetchCallReceiver(input) // Fetch receiver

	t.Logf("found receiver %s", receiver) // Log success
}

// TestCheckAddress - test functionality of CheckAddress() function
func TestCheckAddress(t *testing.T) {
	err := CheckAddress("72.21.215.90") // Attempt to check the address of S3

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log errors
		t.FailNow()           // Panic
	} else if strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: checking addresses requires sudo privileges") // Log warning
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

// TestSHA256 - test functionality of Keccak256() function
func TestKeccak256(t *testing.T) {
	shaVal := Keccak256([]byte("test")) // Create hashed value

	if string(shaVal) == "test" { // Check for errors
		t.Errorf("SHA failed with value %s", shaVal) // Log error
		t.FailNow()                                  // Panic
	}

	t.Logf("hashed value %s", shaVal) // Log success
}
