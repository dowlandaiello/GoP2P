package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/database"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

// handleNewDatabaseCommand - handle execution of handleNewDatabase method (wrapper)
func (term *Terminal) handleNewDatabaseCommand() {
	fmt.Println("attempting to initialize new NodeDatabase") // Log begin

	output, err := term.handleNewDatabase() // Attempt to init new db

	if err != nil { // Check for errors
		fmt.Println("Error: " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleAddNodeCommand - handle execution of handleAddNode method (wrapper)
func (term *Terminal) handleAddNodeCommand(address string) {
	fmt.Println("attempting to add node " + address + " to database") // Log begin

	output, err := term.handleAddNode(address) // Attempt to append

	if err != nil { // Check for errors
		fmt.Println("Error: " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleRemoveNodeCommand - handle execution of handleRemoveNode method (wrapper)
func (term *Terminal) handleRemoveNodeCommand(address string) {
	fmt.Println("attempting to remove node " + address + " from database") // Log begin

	output, err := term.handleRemoveNode(address) // Attempt to remove

	if err != nil { // Check for errors
		fmt.Println("Error: " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleAttachDatabaseCommand - handle execution of handleAttachDatabase method (wrapper)
func (term *Terminal) handleAttachDatabaseCommand() {
	fmt.Println("attempting to attach to NodeDatabase") // Log begin

	output, err := term.handleAttachDatabase() // Attempt to attach to db

	if err != nil { // Check for errors
		fmt.Println("Error: " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleWriteDatabaseToMemoryCommand - handle execution of handleWritDatabaseToMemory method (wrapper)
func (term *Terminal) handleWriteDatabaseToMemoryCommand() {
	fmt.Println("attempting to write database to memory") // Log begin

	output, err := term.handleWriteDatabaseToMemory() // Attempt to write db

	if err != nil { // Check for errors
		fmt.Println("Error: " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleQueryForAddressCommand - handle execution of handleQueryForAddress method (wrapper)
func (term *Terminal) handleQueryForAddressCommand(address string) {
	fmt.Println("attempting to query for address " + address + " in nodedatabase") // Log begin

	output, err := term.handleQueryForAddress(address) // Query

	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		fmt.Println(output)
	}
}

// handleNewDatabase - attempt to initialize new NodeDatabase
func (term *Terminal) handleNewDatabase() (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableType == "Node" { // Verify element is node
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	db, err := database.NewDatabase(&foundNode, "GoP2P_TestNet", common.GoP2PTestNetID, 5) // Attempt to create new database

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable("", db, "NodeDatabase") // Add new database

	err = db.WriteToMemory(foundNode.Environment) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: created new nodedatabase with address " + foundNode.Address, nil // No error occurred, return success
}

// handleAddNode - attempt to append current node to NodeDatabase
func (term *Terminal) handleAddNode(address string) (string, error) {
	if address != "" {
		return term.handleAddSpecificNode(address)
	}

	return term.handleAddCurrentNode()
}

// handleRemoveNode - attempt to remove node from database
func (term *Terminal) handleRemoveNode(address string) (string, error) {
	if address != "" {
		return term.handleRemoveSpecificNode(address)
	}

	return term.handleRemoveCurrentNode()
}

// handleAddSpecificNode - handle execution of addnode method
func (term *Terminal) handleAddSpecificNode(address string) (string, error) {
	foundNode := term.Variables[0].VariableData.(node.Node) // Fetch attached node

	db, err := term.findDatabase() // Attempt to attach to database

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	_, err = db.QueryForAddress(address)

	if err == nil {
		return "", errors.New("node already added to database")
	}

	newNode, err := node.NewNode(address, false) // Attempt to init node with specified address

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = db.AddNode(&newNode) // Attempt to add node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = db.WriteToMemory(foundNode.Environment) // Serialize

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = foundNode.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: added node with address " + address + " to attached node database", nil // Return success
}

// handleRemoveSpecificNode - handle execution of removenode command
func (term *Terminal) handleRemoveSpecificNode(address string) (string, error) {
	foundNode := term.Variables[0].VariableData.(node.Node) // Fetch attached node

	db, err := term.findDatabase()

	if err != nil {
		return "", err
	}

	err = db.RemoveNode(address)

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = db.WriteToMemory(foundNode.Environment) // Serialize

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = foundNode.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: removed node with address " + address + " from attached node database", nil // Return success
}

// handleAddCurrentNode - attempt to add current node to attached node database
func (term *Terminal) handleAddCurrentNode() (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableType == "Node" { // Verify element is node
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	db, err := term.findDatabase() // Attach to database

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	_, qErr := db.QueryForAddress(foundNode.Address) // Check for already existing node

	if qErr != nil { // Check for already existing node
		err := db.AddNode(&foundNode) // Attempt to add node

		if err != nil { // Check for errors
			return "", err // Return found error
		}
	} else { // Node already exists, return error
		return "", errors.New("node already exists in attached database") // Return found error
	}

	err = db.WriteToMemory(term.Variables[0].VariableData.(node.Node).Environment) // Serialize

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = foundNode.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: appended node with address " + foundNode.Address + " to NodeDatabase", nil // No error occurred, return success
}

// handleAddCurrentNode - attempt to add current node to attached node database
func (term *Terminal) handleRemoveCurrentNode() (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableType == "Node" { // Verify element is node
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	db, err := term.findDatabase()

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	_, qErr := db.QueryForAddress(foundNode.Address) // Check for already existing node

	if qErr != nil { // Check for already existing node
		return "", errors.New("node does not exist in attached database") // Node doesn't exist, return error
	}

	err = db.RemoveNode(foundNode.Address) // Attempt to remove node

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = db.WriteToMemory(term.Variables[0].VariableData.(node.Node).Environment) // Serialize

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = foundNode.WriteToMemory(currentDir) // Write node to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: removed node with address " + foundNode.Address + " from NodeDatabase", nil // No error occurred, return success
}

// handleAttachDatabase - handle execution of database reading, write to term mem
func (term *Terminal) handleAttachDatabase() (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableType == "Node" { // Verify element is node
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	db, err := database.ReadDatabaseFromMemory(foundNode.Environment) // Attempt to read database from node environment memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = term.AddVariable("", *db, "NodeDatabase") // Save for persistency

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: attached to nodedatabase with bootstrap address " + (*db.Nodes)[0].Address, nil
}

// handleWritDatabaseToMemory - handle execution of NodeDatabase writeToMemory() method
func (term *Terminal) handleWriteDatabaseToMemory() (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableType == "Node" { // Verify element is node
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	db, err := term.findDatabase() // Attempt to attach to database

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = db.WriteToMemory(foundNode.Environment) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: wrote nodedatabase with address " + foundNode.Address + " to memory", nil // No error occurred, return success
}

// handleQueryForAddress - handle execution of queryforaddress command
func (term *Terminal) handleQueryForAddress(address string) (string, error) {
	db, err := term.findDatabase() // Attach to database

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	index, err := db.QueryForAddress(address) // Query

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: found node with index " + strconv.Itoa(int(index)), nil
}

// findDatabase - attempt to attach to environment database
func (term *Terminal) findDatabase() (*database.NodeDatabase, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableType == "Node" { // Verify element is node
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return &database.NodeDatabase{}, errors.New("node not attached") // Log found error
	}

	db, err := database.ReadDatabaseFromMemory(foundNode.Environment) // Attempt to read database from node environment memory

	if err != nil { // Check for errors
		return &database.NodeDatabase{}, err // Return found error
	}

	return db, nil // No error occurred, return found database
}
