package cli

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/database"
	"github.com/mitsukomegumi/GoP2P/types/handler"
	"github.com/mitsukomegumi/GoP2P/types/node"
	"github.com/mitsukomegumi/GoP2P/upnp"
)

/*
	BEGIN NODE METHODS
*/

// handleNewNodeCommand - handle execution of newnode method
func (term *Terminal) handleNewNodeCommand() {
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond) // Init loading indicator

	s.Prefix = "   "                            // Add line spacing
	s.Suffix = " attempting to create new node" // Add log message

	s.Start() // Start loading indicator

	output, err := term.handleNewNode() // Attempt to initialize new node

	s.Stop() // Stop indicator

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleAttachNodeCommand - attempt to read node at current working directory
func (term *Terminal) handleAttachNodeCommand() {
	fmt.Println("attempting to attach") // Log begin

	output, err := term.handleAttachNode() // Attempt to read node

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleStartHandlerCommand - attempt to start handler on attached node
func (term *Terminal) handleStartHandlerCommand(port int) {
	fmt.Println("attempting to start handler") // Log begin

	output, err := term.handleStartHandler(port) // Attempt to start handler

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleNewNode - handle execution of NewNode() command
func (term *Terminal) handleNewNode() (string, error) {
	node, err := NewNode() // Attempt to create new node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	db, err := database.NewDatabase(node, 5) // Attempt to create new database

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = db.WriteToMemory(node.Environment) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	fmt.Println("\nattempting to write node to memory")

	err = node.WriteToMemory(currentDir) // Write to mem

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	fmt.Println("\n-- SUCCESS -- wrote nodedatabase to environment memory")
	fmt.Println("-- SUCCESS -- wrote node to memory")

	term.AddVariable(db, "NodeDatabase")               // Add new database
	term.AddVariable(*node, "Node")                    // Add new node
	term.AddVariable(*node.Environment, "Environment") // Add new environment

	return "-- SUCCESS -- created node with address " + node.Address, nil // No error occurred, return success
}

// handleAttachNode - handle execution of ReadNode() command
func (term *Terminal) handleAttachNode() (string, error) {
	node, err := AttachNode() // Attempt to read node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	env := node.Environment // Attempt to read environment from memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	db, err := database.ReadDatabaseFromMemory(env) // Attempt to read database from memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(*node, "Node")       // Add node
	term.AddVariable(*env, "Environment") // Add environment
	term.AddVariable(*db, "NodeDatabase") // Add db

	return "-- SUCCESS -- attached to node with address " + node.Address, nil // No error occurred, return success
}

// handleStartHandler - attempt to start handler on node
func (term *Terminal) handleStartHandler(port int) (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Node" { // Verify element is node
			foundNode = term.Variables[x].(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	err := upnp.ForwardPort(3000) // Attempt to forward port

	if err != nil {
		fmt.Println(err.Error())
	}

	ln, err := foundNode.StartListener(port) // Attempt to start handler

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = term.AddVariable(*ln, "Handler") // Attempt to save

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	go handler.StartHandler(&foundNode, ln)

	return "-- SUCCESS -- started handler on port " + strconv.Itoa(port) + " with address " + foundNode.Address, nil // No error occurred, return success
}

/*
	END NODE METHODS
*/
