package common

import "testing"

// TestCheckAddress - test functionality of CheckAddress() function.
func TestCheckAddress(t *testing.T) {
	err := CheckAddress("72.21.215.90") // Attempt to check the address of S3

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log errors
		t.FailNow()           // Panic
	}
}

// TestGetExtIPAddrWithUpNP - test functionality of GetExtIPAddrWithUpNP() function
func TestGetExtIPAddrWithUpNP(t *testing.T) {
	ip, err := GetExtIPAddrWithUpNP() // Attempt to fetch external IP address

	if err != nil { // Check for errors
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
