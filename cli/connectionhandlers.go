package cli

import (
	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/connection"
	"github.com/mitsukomegumi/GoP2P/types/database"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// handleNewConnectionCommand - handle execution of handleNewConnection method
func (term *Terminal) handleNewConnectionCommand(address string, port int) {
	common.Println("attempting to initialize and attempt connection with address " + address) // Log begin

	output, err := term.handleNewConnection(address, port) // Attempt to init, attempt

	if err != nil { // Check for errors
		common.Println("Error: " + err.Error()) // Log error
	} else {
		common.Println(output) // Log success
	}
}

func (term *Terminal) handleNewConnection(address string, port int) (string, error) {
	dbIndex, err := term.QueryType("NodeDatabase") // Fetch db index

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	db := term.Variables[dbIndex].VariableData.(database.NodeDatabase) // Fetch db

	dataIndex, err := term.QueryType("[]byte") // Fetch data index from term

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	data := term.Variables[dataIndex].VariableData.([]byte) // Fetch data

	connectionStackIndex, _ := term.QueryType("[]Event") // Fetch connection stack index

	connectionStack := []connection.Event{} // Create placeholder

	if connectionStackIndex != 0 { // Check for nil index
		connectionStack = term.Variables[connectionStackIndex].VariableData.([]connection.Event) // Fetch connection stack
	}

	sourceNode := term.Variables[0].VariableData.(node.Node) // Fetch sourceNode

	destinationNodeIndex, err := db.QueryForAddress(address) // Fetch destination node index

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	destinationNode := (*db.Nodes)[destinationNodeIndex] // Fetch destination node

	connection, err := connection.NewConnection(&sourceNode, &destinationNode, port, data, "relay", connectionStack) // Init connection

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	_, err = connection.Attempt() // Attempt connection

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "", nil
}
