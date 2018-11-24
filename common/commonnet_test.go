package common

import (
	"crypto/tls"
	"testing"
)

// TestSendBytes - test functionality of SendBytes() method
func TestSendBytes(t *testing.T) {
	err := SendBytes([]byte("test"), "127.0.0.1:443") // Write to address

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("wrote to address 127.0.0.1") // Log success
}

// TestSendBytesWithConnection - test functionality of SendBytesWithConnection() method
func TestSendBytesWithConnection(t *testing.T) {
	connection, err := tls.Dial("tcp", "127.0.0.1:443", GeneralTLSConfig) // Init connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = SendBytesWithConnection(connection, []byte("test")) // Attempt to send bytes

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = connection.Close() // Attempt to close connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}

// TestSendBytesReusable - test functionality of SendBytesReusable() method
func TestSendBytesReusable(t *testing.T) {
	connection, err := SendBytesReusable([]byte("test"), "127.0.0.1:443") // Attempt to send bytes

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = (*connection).Close() // Attempt to close connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}
