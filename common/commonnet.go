package common

import (
	"io"
	"io/ioutil"
	"net"
	"time"
)

// SendBytes - attempt to send specified bytes to given address
func SendBytes(b []byte, address string) error {
	connection, err := net.Dial("tcp", address) // Connect to given address

	if err != nil { // Check for errors
		return err // Return found error
	}

	_, err = connection.Write(b) // Write data to connection

	if err != nil { // Check for errors
		return err // Return found errors
	}

	err = connection.Close() // Close connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}

// SendBytesResult - attempt to send specified bytes to given address, returning result
func SendBytesResult(b []byte, address string) ([]byte, error) {
	connection, err := net.Dial("tcp", address) // Connect to given address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	_, err = connection.Write(b) // Write data to connection

	if err != nil { // Check for errors
		return nil, err // Return found errors
	}

	result, err := ioutil.ReadAll(connection)

	if err != nil { // Check for errors
		return nil, err // Return found errors
	}

	err = connection.Close() // Close connection

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return result, nil // No error occurred, return nil
}

// SendBytesWithConnection - attempt to send specified bytes to given address via given connection
func SendBytesWithConnection(connection *net.Conn, b []byte) error {
	_, err := (*connection).Write(b) // Write to connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}

// SendBytesReusable - attempt to send specified bytes to given address and return created connection
func SendBytesReusable(b []byte, address string) (*net.Conn, error) {
	connection, err := net.Dial("tcp", address) // Connect to given address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	_, err = connection.Write(b) // Write data to connection

	if err != nil { // Check for errors
		return nil, err // Return found errors
	}

	return &connection, nil // No error occurred, return nil
}

// ReadConnectionBytes - attempt to read bytes from connection
func ReadConnectionBytes(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 0, 4096) // Init buffer
	tmp := make([]byte, 256)        // Init temp buffer

	conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // Set read deadline

	for { // Iterate through connection
		n, err := conn.Read(tmp) // Read line

		if err != nil { // Check for errors
			if err != io.EOF { // Check for non-eof
				return []byte{}, err // Return found error
			}

			return buffer, nil // Return read bytes
		}

		buffer = append(buffer, tmp[:n]...) // Append read data
	}
}
