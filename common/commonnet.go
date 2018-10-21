package common

import (
	"bufio"
	"errors"
	"fmt"
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
	b = append(b, byte('\f')) // Append delimiter

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
	b = append(b, byte('\f')) // Append delimiter

	_, err := (*connection).Write(b) // Write to connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}

// SendBytesReusable - attempt to send specified bytes to given address and return created connection
func SendBytesReusable(b []byte, address string) (*net.Conn, error) {
	b = append(b, byte('\f')) // Append delimiter

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
	go func(buffer chan []byte, err chan error) { // Read
		for {
			data := make([]byte, 512) // Init buffer

			_, readErr := conn.Read(data) // Read from connection

			if readErr != nil { // Check for errors
				err <- readErr // Write error

				return // Break
			}

			buffer <- data // Write data
		}
	}(buffer, err)

	ticker := time.Tick(time.Second) // Init time

	for {
		select {
		case data := <-buffer: // Read data from connection
			buffer <- data // Set buffer

			finished <- true // Set finished

			fmt.Println(<-finished) // Log finished
		case <-err: // Check for error
			finished <- true // Set finished

			break // Break loop
		case <-ticker: // Timed out
			err <- errors.New("connection timed out") // Set error

			finished <- true // Set finished
		}
	}
}
