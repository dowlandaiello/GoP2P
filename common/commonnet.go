package common

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strconv"
	"time"
)

/*
	BEGIN EXPORTED METHODS
*/

// SendBytes - attempt to send specified bytes to given address
func SendBytes(b []byte, address string) error {
	connection, err := tls.Dial("tcp", address, GeneralTLSConfig) // Connect to given address

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

// SendBytesNoTLS - send bytes to address without TLS
func SendBytesNoTLS(b []byte, address string) error {
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
	d := net.Dialer{Timeout: 15 * time.Second} // Init dialer with timeout

	connection, err := tls.DialWithDialer(&d, "tcp", address, GeneralTLSConfig) // Connect to given address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	n, err := connection.Write(b) // Write data to connection

	if err != nil { // Check for errors
		return nil, err // Return found error
	} else if n != len(b) { // Check write failed
		return []byte{}, fmt.Errorf("connection write failed: wrote %s bytes of data of %s bytes of data", strconv.Itoa(n), strconv.Itoa(len(b))) // Log connection write failed
	}

	result, err := ReadConnectionWaitAsync(connection) // Read connection

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	err = connection.Close() // Close connection

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return result, nil // No error occurred, return nil
}

// SendBytesAsync - attempt to send specified bytes to given address in an asynchronous manner
func SendBytesAsync(b []byte, address string, finished *[]bool) error {
	connection, err := tls.Dial("tcp", address, GeneralTLSConfig) // Connect to given address

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

	*finished = append(*finished, true) // Append finished

	if finished == nil { // Check for nil
		return errors.New("nil buffer") // Return found error
	}

	return nil // No error occurred, return nil
}

// SendBytesAsyncRoutine - attempt to send specified bytes to given address in an asynchronous, go routine-based manner.
func SendBytesAsyncRoutine(b []byte, address string, finished chan bool) error {
	connection, err := tls.Dial("tcp", address, GeneralTLSConfig) // Connect to given address

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

	finished <- true // Set finished true

	return nil // No error occurred, return nil
}

// SendBytesResultBufferAsync - attempt to send specified bytes to given address in an asynchronous fashion, reading the result into a given buffer
func SendBytesResultBufferAsync(b []byte, buffer *[][]byte, address string) error {
	connection, err := tls.Dial("tcp", address, GeneralTLSConfig) // Connect to given address

	if err != nil { // Check for errors
		return err // Return found error
	}

	_, err = connection.Write(b) // Write data to connection

	if err != nil { // Check for errors
		return err // Return found errors
	}

	result, err := ReadConnectionWaitAsync(connection) // Read connection

	if err != nil { // Check for errors
		return err // Return found errors
	}

	err = connection.Close() // Close connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	*buffer = append(*buffer, result) // Append result

	if buffer == nil { // Check for nil buffer
		return errors.New("nil buffer") // Return found error
	}

	return nil // No error occurred, return nil
}

// SendBytesWithConnection - attempt to send specified bytes to given address via given connection
func SendBytesWithConnection(connection *tls.Conn, b []byte) error {
	_, err := (*connection).Write(b) // Write to connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}

// SendBytesReusable - attempt to send specified bytes to given address and return created connection
func SendBytesReusable(b []byte, address string) (*tls.Conn, error) {
	connection, err := tls.Dial("tcp", address, GeneralTLSConfig) // Connect to given address

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	_, err = connection.Write(b) // Write data to connection

	if err != nil { // Check for errors
		return nil, err // Return found errors
	}

	return connection, nil // No error occurred, return nil
}

// ReadConnectionDelim - attempt to read connection until occurrence of standard GoP2P connection delimiter
func ReadConnectionDelim(conn *tls.Conn) ([]byte, error) {
	reader := bufio.NewReader(conn) // Initialize reader

	data, err := reader.ReadBytes(ConnectionDelimiter) // Read until delimiter

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	return data, nil // Return read data
}

// ReadConnectionAsync - attempt to read entirety of specified connection in an asynchronous fashion, returning data byte value
func ReadConnectionAsync(conn *tls.Conn, buffer chan []byte, finished chan bool, err chan error) {
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
func ReadConnectionWaitAsync(conn *tls.Conn) ([]byte, error) {
	data := make(chan []byte) // Init buffer
	err := make(chan error)   // Init error buffer

	conn.SetReadDeadline(time.Now().Add(4 * time.Second)) // Set read deadline

	go func(data chan []byte, err chan error) {
		reads := 0 // Init reads buffer

		for {
			reads++ // Increment read

			var buffer bytes.Buffer // Init buffer

			readData, readErr := io.Copy(&buffer, conn) // Read connection

			if readErr != nil && readErr != io.EOF && reads > 3 { // Check for errors
				err <- readErr // Write read error
			} else if readData == 0 { // Check for nil readData
				continue // Continue
			}

			data <- buffer.Bytes() // Write read data
		}
	}(data, err)

	ticker := time.Tick(5 * time.Second) // Init ticker

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

// ReadConnectionWaitAsyncNoTLS - attempt to read from connection in an asynchronous fashion, after waiting for peer to write
func ReadConnectionWaitAsyncNoTLS(conn net.Conn) ([]byte, error) {
	data := make(chan []byte) // Init buffer
	err := make(chan error)   // Init error buffer

	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) // Set read deadline

	go func(data chan []byte, err chan error) {
		reads := 0 // Init reads buffer

		for {
			reads++ // Increment read

			var buffer bytes.Buffer // Init buffer

			readData, readErr := io.Copy(&buffer, conn) // Read connection

			if readErr != nil && readErr != io.EOF && reads > 3 { // Check for errors
				err <- readErr // Write read error
			} else if readData == 0 { // Check for nil readData
				continue // Continue
			}

			data <- buffer.Bytes() // Write read data
		}
	}(data, err)

	ticker := time.Tick(3 * time.Second) // Init ticker

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
