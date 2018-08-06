package handler

import (
	"errors"
	"io/ioutil"
	"net"
	"reflect"

	"github.com/mitsukomegumi/GoP2P/types/connection"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// StartHandler - attempt to accept and handle requests on given listener
func StartHandler(node *node.Node, ln *net.Listener) error {

	if reflect.ValueOf(node).IsNil() || node.Address == "" || reflect.ValueOf(ln).IsNil() { // Check for nil parameters
		return errors.New("invalid parameters") // Return error
	}

	for {
		conn, err := (*ln).Accept() // Accept connection

		if err == nil { // Check for errors
			go handleConnection(node, conn) // Handle connection
		}
	}
}

func handleConnection(node *node.Node, conn net.Conn) error {
	data, err := ioutil.ReadAll(conn) // Attempt to read from connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	_, err = connection.FromBytes(data) // Attempt to decode data

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}
