package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

/*
	BEGIN ENVIRONMENT METHODS
*/

// handleAttachEnvironmentCommand - attempt to read environment at current working directory
func (term *Terminal) handleAttachEnvironmentCommand() {
	common.Println("attempting to attach") // Log begin

	output, err := term.handleAttachEnvironment() // Attempt to read env

	if err != nil { // Check for errors
		common.Println("Error: " + err.Error()) // Log error
	} else {
		common.Println(output) // Log success
	}
}

// handleNewEnvironmentCommand - handle execution of handleNewEnvironment method (wrapper)
func (term *Terminal) handleNewEnvironmentCommand() {
	common.Println("attempting to initialize new environment") // Log begin

	output, err := term.handleNewEnvironment() // Attempt to init new environment

	if err != nil { // Check for errors
		common.Println("Error: " + err.Error()) // Log error
	} else {
		common.Println(output) // Log success
	}
}

func (term *Terminal) handleNewVariableCommand(variableType string, variableDir string, variableData string, replaceExisting bool) {
	common.Println("attempting to add new variable") // Log begin

	output, err := term.handleNewVariable(variableType, variableDir, variableData, replaceExisting) // Execute command

	if err != nil { // Check for errors
		common.Println("Error: " + err.Error()) // Log error
	} else {
		common.Println(output) // Log success
	}
}

func (term *Terminal) handleWriteToMemoryCommand() {
	common.Println("attempting to environment to memory") // Log begin

	output, err := term.handleWriteToMemory() // Attempt to write env

	if err != nil { // Check for errors
		common.Println("Error: " + err.Error()) // Log error
	} else {
		common.Println(output) // Log success
	}
}

// handleQueryTypeCommand - handle execution of handleQueryType method (wrapper)
func (term *Terminal) handleQueryTypeCommand(queryType string) {
	common.Println("querying type " + queryType) // Log begin

	output, err := term.handleQueryType(queryType) // Attempt to query for type

	if err != nil { // Check for errors
		common.Println("Error: " + err.Error()) // Log error
	} else {
		common.Println(output) // Log success
	}
}

// handleQueryValueCommand - handle execution of handleQueryValue method (wrapper)
func (term *Terminal) handleQueryValueCommand(queryValue string) {
	common.Println("querying value " + queryValue) // Log begin

	output, err := term.handleQueryValue(queryValue) // Attempt to query for value

	if err != nil { // Check for errors
		common.Println("Error: " + err.Error()) // Log error
	} else {
		common.Println(output) // Log success
	}
}

// handleNewEnvironment - attempt to initialize new environment
func (term *Terminal) handleNewEnvironment() (string, error) {
	foundNode := node.Node{}

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableData == "Node" { // Verify element is node
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	if foundNode.Address == "" { // Check for errors
		return "", errors.New("node not attached") // Log found error
	}

	env, err := environment.NewEnvironment() // Attempt to create new environment

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable("", *env, "Environment") // Add new environment

	foundNode.Environment = env // Set environment

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = foundNode.WriteToMemory(currentDir) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: created new environment", nil // No error occurred, return success
}

// handleNewVariable - attempt to init and append new variable to environment variables list
func (term *Terminal) handleNewVariable(variableType string, dir string, variableData string, replaceExisting bool) (string, error) {
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond) // Init loading indicator

	s.Prefix = "   "                                    // Add line spacing
	s.Suffix = " attempting to initialize new variable" // Add log message

	s.Start() // Start loading indicator

	var data []byte // Init buffer

	var err error // Init error

	if variableData == "" { // Check for data
		file, err := os.Open(dir) // Attempt to open file at specified directory

		if err != nil { // Check for errors
			return "", err // Return found errors
		}

		data, err = ioutil.ReadAll(file) // Attempt to read file

		if err != nil { // Check for errors
			return "", err // Return found error
		}
	}

	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableData == "Node" { // Verify element is environment
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	variable := &environment.Variable{}

	if variableData == "" {
		variable, err = environment.NewVariable(variableType, data) // Init variable
	} else {
		variable, err = environment.NewVariable(variableType, variableData) // Init variable
	}

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = (*foundNode.Environment).AddVariable(variable, replaceExisting) // Append

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = foundNode.WriteToMemory(currentDir) // Write node to mem

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	s.Stop() // Stop loading indicator

	return "Success: initialized and added variable with type " + variable.VariableType, err
}

// handleQueryType - attempt to query for specified type in environment
func (term *Terminal) handleQueryType(queryType string) (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableData == "Node" { // Verify element is node
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	value, err := foundNode.Environment.QueryType(queryType) // Attempt to query for type

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	strVal, err := common.SerializeToString(*value) // Serialize response

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: found variable with type " + strVal, nil // Return response
}

// handleQueryValue - attempt to query for specified value in environment
func (term *Terminal) handleQueryValue(queryValue string) (string, error) {
	foundNode := node.Node{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableData == "Node" { // Verify element is environment
			foundNode = term.Variables[x].VariableData.(node.Node) // Set to value

			break
		}
	}

	value, err := foundNode.Environment.QueryValue(queryValue) // Attempt to query for value

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	strVal, err := common.SerializeToString(*value) // Serialize response

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: found variable with value " + strVal, nil // Return response
}

// handleAttachEnvironment - handle execution of ReadEnvironment() command
func (term *Terminal) handleAttachEnvironment() (string, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	node, err := node.ReadNodeFromMemory(currentDir) // Attempt to read environment

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable("", *node.Environment, "Environment") // Add env to variables

	return "Success: attached to environment", nil // No error occurred, return success
}

// handleWriteToMemory - handle execution of WriteToMemory() method
func (term *Terminal) handleWriteToMemory() (string, error) {
	foundEnvironment := environment.Environment{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.Variables[x].VariableData == "Environment" { // Verify element is environment
			foundEnvironment = term.Variables[x].VariableData.(environment.Environment) // Set to value

			break
		}
	}

	currentDir, err := common.GetCurrentDir() // Attempt to fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = foundEnvironment.WriteToMemory(currentDir) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "Success: wrote environment to memory", nil // Return success
}

func handleNilVarParams() (string, string, string, bool) {
	var variableType string    // Init buffer
	var variableDir string     // Init buffer
	var variableData string    // Init buffer
	var replaceExisting string // Init buffer

	common.Print("\nvariable type: ")
	fmt.Scanln(&variableType)

	common.Print("\nvariable data directory (optional): ")
	fmt.Scanln(&variableDir)

	common.Print("\nvariable data (optional): ")
	fmt.Scanln(&variableData)

	common.Print("\nreplace existing directory: ")
	fmt.Scanln(&replaceExisting)

	boolVal, _ := strconv.ParseBool(replaceExisting)

	return variableType, variableDir, variableData, boolVal
}

func handleVarParams(command string) (string, string, string, bool) {
	params := strings.Split(strings.Split(strings.Split(command, "(")[1], ")")[0], ", ") // Split command

	boolVal, _ := strconv.ParseBool(params[len(params)-1]) // Parse replace existing

	if strings.Contains(command, "/") || strings.Contains(command, "\\") {
		return params[0], params[1], "", boolVal // Return values
	}

	return params[0], "", params[1], boolVal // Return values
}

/*
	END ENVIRONMENT METHODS
*/
