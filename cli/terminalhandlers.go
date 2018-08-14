package cli

import (
	"fmt"
	"strconv"
	"strings"
)

func (term *Terminal) handleCommand(command string) {
	switch { // Iterate through possible commands
	case strings.Contains(strings.ToLower(command), "upnp."): // Account for UpNP methods
		term.handleUpNP(command)
	case strings.Contains(strings.ToLower(command), "node."): // Account for node methods
		term.handleNode(command)
	case strings.Contains(strings.ToLower(command), "environment."): // Account for environment methods
		term.handleEnvironment(command)
	case strings.Contains(strings.ToLower(command), "database."): // Account for nodedatabase methods
		term.handleDatabase(command)
	}
}

/*
	BEGIN METHOD ROUTING
*/

func (term *Terminal) handleUpNP(command string) {
	switch {
	case strings.Contains(strings.ToLower(command), "forwardport"): // Account for forwardport command
		intVal, _ := strconv.Atoi(strings.Split(strings.Split(command, "(")[1], ")")[0]) // Attempt to fetch port from command

		if intVal == 0 {
			intVal = handleZeroPort() // Fetch port
		}

		term.handleForwardPortCommand(intVal) // Forward port
	case strings.Contains(strings.ToLower(command), "removeportforward"): // Account for removeportforward command
		intVal, _ := strconv.Atoi(strings.Split(strings.Split(command, "(")[1], ")")[0]) // Attempt to fetch port from command

		if intVal == 0 {
			intVal = handleZeroPort() // Fetch port
		}

		term.handleRemoveForwardPortCommand(intVal) // Remove port forwarding
	}
}

func (term *Terminal) handleNode(command string) {
	switch {
	case strings.Contains(strings.ToLower(command), "newnode"): // Account for newnode command
		term.handleNewNodeCommand()
	case strings.Contains(strings.ToLower(command), "attach"): // Account for readnode command
		term.handleAttachNodeCommand()
	case strings.Contains(strings.ToLower(command), "starthandler"):
		intVal, _ := strconv.Atoi(strings.Split(strings.Split(command, "(")[1], ")")[0]) // Attempt to fetch port from command

		if intVal == 0 {
			intVal = handleZeroPort() // Fetch port
		}

		term.handleStartHandlerCommand(intVal) // Start handler command execution
	}
}

func (term *Terminal) handleEnvironment(command string) {
	switch {
	case strings.Contains(strings.ToLower(command), "newenvironment"): // Account for environment initializer
		term.handleNewEnvironmentCommand() // Execute command
	case strings.Contains(strings.ToLower(command), "newvariable"): // Account for newvariable method
		var variableType, variableDir, variableData string
		var replaceExisting bool

		if !strings.Contains(command, ",") {
			variableType, variableDir, variableData, replaceExisting = handleNilVarParams()
		} else {
			variableType, variableDir, variableData, replaceExisting = handleVarParams(command)
		}

		term.handleNewVariableCommand(variableType, variableDir, variableData, replaceExisting)
	case strings.Contains(strings.ToLower(command), "querytype"): // Account for querytype method
		queryType := strings.Split(strings.Split(command, "(")[1], ")")[0] // Fetch value from command

		term.handleQueryTypeCommand(queryType) // Execute command
	case strings.Contains(strings.ToLower(command), "queryvalue"): // Account for queryvalue method
		queryValue := strings.Split(strings.Split(command, "(")[1], ")")[0] // Fetch value from command

		term.handleQueryValueCommand(queryValue) // Execute command
	case strings.Contains(strings.ToLower(command), "attach"): // Account for attach method
		term.handleAttachEnvironmentCommand() // Execute command
	case strings.Contains(strings.ToLower(command), "writetomemory"): // Account for i/o methods
		term.handleWriteToMemoryCommand() // Execute command
	}
}

func (term *Terminal) handleDatabase(command string) {
	switch {
	case strings.Contains(strings.ToLower(command), "newdatabase"): // Account for init method
		term.handleNewDatabaseCommand() // Execute command
	case strings.Contains(strings.ToLower(command), "writetomemory"): // Account for i/o methods
		term.handleWriteDatabaseToMemoryCommand() // Execute command
	case strings.Contains(strings.ToLower(command), "attach"): // Account for attach method
		term.handleAttachDatabaseCommand() // Execute command
	case strings.Contains(strings.ToLower(command), "addnode"): // Account for AddNode method
		address := strings.Split(strings.Split(command, "(")[1], ")")[0]

		term.handleAddNodeCommand(address) // Execute command
	}
}

/*
	END METHOD ROUTING
*/

/*
	BEGIN GENERAL METHODS
*/

// handleZeroPort - handle circumstance in which user has not specified a port
func handleZeroPort() int {
	var input string // Init buffer

	fmt.Print("\nport: ")
	fmt.Scanln(&input)

	intVal, _ := strconv.Atoi(input) // Convert to int

	return intVal // Return result
}

/*
	END GENERAL METHODS
*/
