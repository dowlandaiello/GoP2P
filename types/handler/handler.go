package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/connection"
	"github.com/mitsukomegumi/GoP2P/types/environment"
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
	fmt.Printf("-- CONNECTION -- address: %s", conn.RemoteAddr().String())

	data, err := ioutil.ReadAll(conn) // Attempt to read from connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	readConnection, err := connection.FromBytes(data) // Attempt to decode data

	if err != nil { // Check for errors
		return err // Return found error
	}

	fmt.Println("-- CONNECTION " + conn.RemoteAddr().String() + " -- attempted to read " + strconv.Itoa(len(data)) + " byte of data.")

	if len(readConnection.ConnectionStack) == 0 { // Check if event stack exists
		val, err := handleSingular(node, readConnection) // Handle singular event

		if err != nil { // Check for errors
			return err // Return found error
		}

		conn.Write(val) // Write success

		return nil // No error occurred, return nil
	}

	val, err := handleStack(node, readConnection) // Attempt to handle stack

	if err != nil { // Check for errors
		return err // Return found error
	}

	instancedResponse := connection.Response{Val: val} // Create response instance for byte serialization

	serializedResponse, err := common.SerializeToBytes(instancedResponse) // Serialize

	if err != nil { // Check for errors
		return err // Return found error
	}

	conn.Write(serializedResponse) // Write success

	return nil // Attempt to handle stack
}

func handleSingular(node *node.Node, connection *connection.Connection) ([]byte, error) {
	variable, err := environment.NewVariable("byte[]", connection) // Init variable to hold connection data

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	varByteVal, err := common.SerializeToBytes(variable) // Serialize

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return varByteVal, node.Environment.AddVariable(variable) // Attempt to add variable to environment, return variable value as byte
}

func handleStack(node *node.Node, connection *connection.Connection) ([][]byte, error) {
	for x := 0; x != len(connection.ConnectionStack); x++ { // Iterate through stack
		handleCommand(node, &connection.ConnectionStack[x]) // Attempt to handle command
	}

	return nil, nil // No error occurred, return nil
}

func handleCommand(node *node.Node, event *connection.Event) ([]byte, error) {
	switch {
	case strings.Contains(event.Command.Command, "NewVariable"):
		return handleNewVariable(node, event) // Attempt command
	case strings.Contains(event.Command.Command, "QueryValue"):
		return handleQueryValue(node, event) // Attempt command
	case strings.Contains(event.Command.Command, "QueryType"):
		return handleQueryType(node, event) // Attempt command
	case strings.Contains(event.Command.Command, "AddVariable"):
		return handleAddVariable(node, event) // Attempt command
	default:
		return nil, nil // Return nil value
	}
}

func handleNewVariable(node *node.Node, event *connection.Event) ([]byte, error) {
	variableType := event.Command.ModifierSet.Type // Attempt to fetch variable type from command

	variable, err := environment.NewVariable(variableType, event.Command.ModifierSet.Value) // Attempt to create new variable

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	serializedValue, err := common.SerializeToBytes(variable) // Attempt to serialize new variable

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return serializedValue, nil // Return serialized value
}

func handleQueryValue(node *node.Node, event *connection.Event) ([]byte, error) {
	return nil, nil // Return result
}

func handleQueryType(node *node.Node, event *connection.Event) ([]byte, error) {
	return nil, nil // Return result
}

func handleAddVariable(node *node.Node, event *connection.Event) ([]byte, error) {
	return nil, nil // Return result
}
