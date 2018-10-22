package common

import (
	"bufio"
	"io/ioutil"
	"net"
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

	result, err := ioutil.ReadAll(connection) // Read connection

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

// ReadConnectionDelim - attempt to read connection until occurrence of standard GoP2P connection delimiter
func ReadConnectionDelim(conn net.Conn) ([]byte, error) {
	reader := bufio.NewReader(conn) // Initialize reader

	data, err := reader.ReadBytes(ConnectionDelimiter) // Read until delimiter

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	return data, nil // Return read data
}

// ReadConnectionAsync - attempt to read entirety of specified connection in an asynchronous fashion, returning data byte value
func ReadConnectionAsync(conn net.Conn, buffer chan []byte, finished chan bool, err chan error) {
	data, readErr := ioutil.ReadAll(conn) // Read connection

	if readErr != nil { // Check for errors
		err <- readErr   // Set error
		finished <- true // Set finished

		return
	}

	buffer <- data // Set read data

	finished <- true // Set finished

	return
}
