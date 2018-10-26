package common

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"time"
)

/*
	BEGIN EXPORTED METHODS
*/

// SendBytes - attempt to send specified bytes to given address
func SendBytes(b []byte, address string) error {
	b = bytes.Trim(b, "\x00") // Trim line end

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
	b = bytes.Trim(b, "\x00") // Trim line end

	connection, err := net.Dial("tcp", address) // Connect to given address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	_, err = connection.Write(b) // Write data to connection

	if err != nil { // Check for errors
		return nil, err // Return found errors
	}

	result, err := ReadConnectionWaitAsync(connection) // Read connection

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
	b = bytes.Trim(b, "\x00") // Trim line end

	_, err := (*connection).Write(b) // Write to connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}

// SendBytesReusable - attempt to send specified bytes to given address and return created connection
func SendBytesReusable(b []byte, address string) (*net.Conn, error) {
	b = bytes.Trim(b, "\x00") // Trim line end

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

// ReadConnectionWaitAsync - attempt to read from connection in an asynchronous fashion, after waiting for peer to write
func ReadConnectionWaitAsync(conn net.Conn) ([]byte, error) {
	data := make(chan []byte) // Init buffer
	err := make(chan error)   // Init error buffer

	go func(data chan []byte, err chan error) {
		for {
			readData := make([]byte, 2048) // Init read buffer

			_, readErr := conn.Read(readData) // Read into buffer

			if readErr != nil { // Check for errors
				err <- readErr // Write found error

				return // Return
			}

			data <- readData // Write read data
		}
	}(data, err)

	ticker := time.Tick(time.Second) // Init ticker

	for { // Continuously read from connection
		select {
		case readData := <-data: // Read data from connection
			return readData, nil // Return read data
		case readErr := <-err: // Error on read
			return []byte{}, readErr // Return error
		case <-ticker: // Timed out
			return []byte{}, errors.New("timed out") // Return timed out error
		}
	}
}

/*
	END EXPORTED METHODS
*/
