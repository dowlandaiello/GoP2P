package cli

import (
	"fmt"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
)

/*
	BEGIN ENVIRONMENT METHODS
*/

// handleAttachEnvironmentCommand - attempt to read environment at current working directory
func (term *Terminal) handleAttachEnvironmentCommand() {
	fmt.Println("attempting to attach") // Log begin

	output, err := term.handleAttachEnvironment() // Attempt to read env

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleNewEnvironmentCommand - handle execution of handleNewEnvironment method (wrapper)
func (term *Terminal) handleNewEnvironmentCommand() {
	fmt.Println("attempting to initialize new environment") // Log begin

	output, err := term.handleNewEnvironment() // Attempt to init new environment

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

func (term *Terminal) handleWriteToMemoryCommand() {
	fmt.Println("attempting to environment to memory") // Log begin

	output, err := term.handleWriteToMemory() // Attempt to write env

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleQueryTypeCommand - handle execution of handleQueryType method (wrapper)
func (term *Terminal) handleQueryTypeCommand(queryType string) {
	fmt.Println("querying type " + queryType) // Log begin

	output, err := term.handleQueryType(queryType) // Attempt to query for type

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleQueryValueCommand - handle execution of handleQueryValue method (wrapper)
func (term *Terminal) handleQueryValueCommand(queryValue string) {
	fmt.Println("querying value " + queryValue) // Log begin

	output, err := term.handleQueryValue(queryValue) // Attempt to query for value

	if err != nil { // Check for errors
		fmt.Println("-- ERROR -- " + err.Error()) // Log error
	} else {
		fmt.Println(output) // Log success
	}
}

// handleNewEnvironment - attempt to initialize new environment
func (term *Terminal) handleNewEnvironment() (string, error) {
	env, err := environment.NewEnvironment() // Attempt to create new environment

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(*env, "Environment") // Add new environment

	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	err = env.WriteToMemory(currentDir) // Attempt to write to memory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- created new environment", nil // No error occurred, return success
}

// handleQueryType - attempt to query for specified type in environment
func (term *Terminal) handleQueryType(queryType string) (string, error) {
	foundEnvironment := environment.Environment{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Environment" { // Verify element is environment
			foundEnvironment = term.Variables[x].(environment.Environment) // Set to value

			break
		}
	}

	value, err := foundEnvironment.QueryType(queryType) // Attempt to query for type

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	strVal, err := common.SerializeToString(*value) // Serialize response

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- found variable with type " + strVal, nil // Return response
}

// handleQueryValue - attempt to query for specified value in environment
func (term *Terminal) handleQueryValue(queryValue string) (string, error) {
	foundEnvironment := environment.Environment{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Environment" { // Verify element is environment
			foundEnvironment = term.Variables[x].(environment.Environment) // Set to value

			break
		}
	}

	value, err := foundEnvironment.QueryValue(queryValue) // Attempt to query for value

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	strVal, err := common.SerializeToString(*value) // Serialize response

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return "-- SUCCESS -- found variable with value " + strVal, nil // Return response
}

// handleAttachEnvironment - handle execution of ReadEnvironment() command
func (term *Terminal) handleAttachEnvironment() (string, error) {
	currentDir, err := common.GetCurrentDir() // Fetch working directory

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	env, err := environment.ReadEnvironmentFromMemory(currentDir) // Attempt to read environment

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	term.AddVariable(*env, "Environment") // Add env to variables

	return "-- SUCCESS -- attached to environment", nil // No error occurred, return success
}

// handleWriteToMemory - handle execution of WriteToMemory() method
func (term *Terminal) handleWriteToMemory() (string, error) {
	foundEnvironment := environment.Environment{} // Create placeholder

	for x := 0; x != len(term.Variables); x++ { // Iterate through array
		if term.VariableTypes[x] == "Environment" { // Verify element is environment
			foundEnvironment = term.Variables[x].(environment.Environment) // Set to value

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

	return "-- SUCCESS -- wrote environment to memory", nil // Return success
}

/*
	END ENVIRONMENT METHODS
*/
