package cli

import (
	"errors"
	"fmt"

	"github.com/mitsukomegumi/GoP2P/types/database"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// handleNewDatabaseCommand - handle execution of handleNewDatabase method (wrapper)
func (term *Terminal) handleNewDatabaseCommand() {
	fmt.Println("attempting to initialize new NodeDatabase") // Log begin

	output, err := term.handleNewDatabase() // Attempt to init new db

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleAttachDatabaseCommand - handle execution of handleAttachDatabase method (wrapper)
func (term *Terminal) handleAttachDatabaseCommand() {
	fmt.Println("attempting to attach to NodeDatabase") // Log begin

	output, err := term.handleAttachDatabase() // Attempt to attach to db

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleWriteDatabaseToMemoryCommand - handle execution of handleWritDatabaseToMemory method (wrapper)
func (term *Terminal) handleWriteDatabaseToMemoryCommand() {
	fmt.Println("attempting to write database to memory") // Log begin

	output, err := term.handleWriteDatabaseToMemory() // Attempt to write db

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleNewDatabase - attempt to initialize new NodeDatabase
func (term *Terminal) handleNewDatabase() (string, error) {
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

	db, err := database.NewDatabase(&foundNode, 5) // Attempt to create new database

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(db, "NodeDatabase") // Add new database

	err = db.WriteToMemory(foundNode.Environment) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- created new nodedatabase with address " + foundNode.Address, nil // No error occurred, return success
}

// handleAttachDatabase - handle execution of database reading, write to term mem
func (term *Terminal) handleAttachDatabase() (string, error) {
	return "", nil
}

// handleWritDatabaseToMemory - handle execution of NodeDatabase writeToMemory() method
func (term *Terminal) handleWriteDatabaseToMemory() (string, error) {
	foundNode := node.Node{}           // Create placeholder
	foundDb := database.NodeDatabase{} // Create placeholder

	emptyNode := node.Node{}
	emptyDb := database.NodeDatabase{}

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Node" { // Verify element is node
			foundNode = term.Variables[x].(node.Node) // Set to value

			if foundDb != emptyDb { // Check for valid db
				break
			}
		} else if term.VariableTypes[x] == "NodeDatabase" { // Verify element is NodeDatabase
			foundDb = term.Variables[x].(database.NodeDatabase) // Set to value

			if foundNode != emptyNode { // Check for valid node
				break
			}
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	err := foundDb.WriteToMemory(foundNode.Environment) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- wrote nodedatabase with address " + foundNode.Address + " to memory", nil // No error occurred, return success
}
