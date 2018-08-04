package common

import "testing"

// TestSendBytes - test functionality of SendBytes() method
func TestSendBytes(t *testing.T) {
	err := SendBytes([]byte("test"), "1.1.1.1:80") // Write to address

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("wrote to address 1.1.1.1") // Log success
}
