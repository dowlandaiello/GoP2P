package handler

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/connection"
	"github.com/mitsukomegumi/GoP2P/types/database"
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

// handleConnection - attempt to fetch connection metadata, handle it respectively (stack or singular)
func handleConnection(node *node.Node, conn net.Conn) error {
	data, err := common.ReadConnectionWaitAsync(conn) // Read entire connection

	if err != nil { // Check for errors
		return err // Return found error
	}

	if strings.Contains(string(data), "messagetype") { // Check for network message
		return handleLogNetworkMessage(data) // Handle network message
	}

	fmt.Printf("\n-- CONNECTION -- incoming connection from address: %s with data %s", conn.RemoteAddr().String(), string(data)) // Log connection

	readConnection, err := connection.FromBytes(data) // Attempt to decode data

	if err != nil { // Check for errors
		return err // Return found error
	}

	fmt.Println("\n\n-- CONNECTION " + conn.RemoteAddr().String() + " -- attempted to read " + strconv.Itoa(len(data)) + " byte of data.") // Log read connection

	if len(readConnection.ConnectionStack) == 0 { // Check if event stack exists
		val, isMessage, err := handleSingular(node, readConnection) // Handle singular event

		if err != nil { // Check for errors
			return err // Return found error
		}

		serializedResponse, err := common.SerializeToBytes(connection.Response{Val: [][]byte{val}}) // Attempt to serialize response

		if err != nil { // Check for errors
			return err // Return found error
		}

		if isMessage == true {
			handleLogNetworkMessage(val) // Handle network message
		} else {
			fmt.Println("\n-- CONNECTION " + conn.RemoteAddr().String() + " -- responding with data " + common.SafeSlice(serializedResponse) + "...") // Log response
		}

		conn.Write(serializedResponse) // Write success

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

	fmt.Println("\n-- CONNECTION " + readConnection.InitializationNode.Address + " -- responding with data " + common.SafeSlice(serializedResponse) + "...") // Log response

	conn.Write(serializedResponse) // Write success

	return nil // Attempt to handle stack
}

// handleSingular - no stack present in found connection, write variable with connection data
func handleSingular(node *node.Node, connection *connection.Connection) ([]byte, bool, error) {
	db, err := database.FromBytes(connection.Data) // Attempt to read db

	if err == nil { // Check for success
		err = db.WriteToMemory(node.Environment) // Write db to memory

		if err != nil { // Check for errors
			return nil, false, err // Return found error
		}

		result, err := common.SerializeToBytes(*db) // Attempt to serialize

		if err != nil { // Check for errors
			return nil, false, err // Return found error
		}

		return result, false, nil // Attempt to serialize
	}

	result, err := handleNetworkMessage(node, connection) // Attempt to decode message

	if err == nil { // Check for success
		return result, true, nil // Return result
	}

	variable, err := environment.NewVariable("Connection", connection) // Init variable to hold connection data

	if err != nil { // Check for errors
		return nil, false, err // Return found error
	}

	varByteVal, err := common.SerializeToBytes(variable) // Serialize

	if err != nil { // Check for errors
		return nil, false, err // Return found error
	}

	return varByteVal, false, node.Environment.AddVariable(variable, false) // Attempt to add variable to environment, return variable value as byte
}

// handleLogNetworkMessage - handle logging of a network message
func handleLogNetworkMessage(b []byte) error {
	message, err := database.MessageFromBytes(b) // Fetch message from connection data

	if err != nil { // Check for errors
		return err // Return found error
	}

	red := color.New(color.FgRed)       // Init red writer
	yellow := color.New(color.FgYellow) // Init yellow writer
	cyan := color.New(color.FgCyan)     // Init cyan writer

	switch message.Priority { // Account for different message priorities
	case 0: // Check for normal message
		fmt.Printf("\n== NETWORK MESSAGE (%s) == %s", message.Type, message.Message) // Log response
	case 1: // Check for critical message
		red.Printf("\n== CRITICAL NETWORK MESSAGE (%s) == %s", strings.ToUpper(message.Type), message.Message) // Log response
	case 2: // Check for warning message
		yellow.Printf("\n== NETWORK MESSAGE (%s) == %s", message.Type, message.Message) // Log response
	case 3: // Check for update/info message
		cyan.Printf("\n== NETWORK MESSAGE (%s) == %s", message.Type, message.Message) // Log response
	}

	return nil // No error occurred, return nil
}

// handleNetworkMessage - handle received network message
func handleNetworkMessage(node *node.Node, connection *connection.Connection) ([]byte, error) {
	message, err := database.MessageFromBytes(connection.Data) // Fetch message from connection data

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	variable, err := environment.NewVariable(fmt.Sprintf("%sNetworkMessage", message.Network), *message) // Init variable

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	err = node.Environment.AddVariable(variable, false) // Attempt to add variable to environment

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	return message.ToBytes() // Return message value
}

// handleStack - found connection with stack, iterate through and handle each command
func handleStack(node *node.Node, connection *connection.Connection) ([][]byte, error) {
	responses := [][]byte{} // Create placeholder

	for x := 0; x != len(connection.ConnectionStack); x++ { // Iterate through stack
		val, _ := handleCommand(node, &connection.ConnectionStack[x]) // Attempt to handle command

		responses = append(responses, val) // Append response
	}

	if len(responses) == 0 {
		return nil, errors.New("nil response")
	}

	return responses, nil // No error occurred, return nil
}

func handleCommand(node *node.Node, event *connection.Event) ([]byte, error) {
	refreshedNode, err := refreshNode() // Refresh node

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	*node = *refreshedNode // Reset refreshed node

	switch event.Command.Command { // Check for commands
	case "NewVariable":
		return handleNewVariable(node, event) // Attempt command
	case "QueryValue":
		return handleQueryValue(node, event) // Attempt command
	case "QueryType":
		return handleQueryType(node, event) // Attempt command
	case "AddVariable":
		return handleAddVariable(node, event) // Attempt command
	default:
		return nil, errors.New("invalid command " + event.Command.Command) // Return nil value
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
	variable, err := node.Environment.QueryValue(event.Command.ModifierSet.Value.(string)) // Attempt to query for value

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	serializedValue, err := common.SerializeToBytes(variable) // Attempt to serialize new variable

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return serializedValue, nil // Return serialized value
}

func handleQueryType(node *node.Node, event *connection.Event) ([]byte, error) {
	variable, err := node.Environment.QueryType(event.Command.ModifierSet.Type) // Attempt to query for value

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	serializedValue, err := common.SerializeToBytes(variable) // Attempt to serialize new variable

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return serializedValue, nil // Return serialized value
}

func handleAddVariable(node *node.Node, event *connection.Event) ([]byte, error) {
	variable := event.Command.ModifierSet.Variable // Attempt to fetch variable from command

	if reflect.ValueOf(variable).IsNil() { // Check for errors
		return nil, errors.New("nil variable") // Return found nil variable
	}

	err := node.Environment.AddVariable(variable, false) // Attempt to add found variable to environment

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	serializedValue, err := common.SerializeToBytes(variable) // Attempt to serialize variable

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return serializedValue, nil // Return serialized value
}

func refreshNode() (*node.Node, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return &node.Node{}, err // return found error
	}

	readNode, err := node.ReadNodeFromMemory(currentDir) // Read node from working dir

	if err != nil { // Check for errors
		return &node.Node{}, err // return found error
	}

	return readNode, nil // Return found node
}
