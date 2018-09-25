package cli

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitsukomegumi/GoP2P/common"
	handlerProto "github.com/mitsukomegumi/GoP2P/rpc/proto/handler"
	nodeProto "github.com/mitsukomegumi/GoP2P/rpc/proto/node"
)

// Terminal - absctract container holding set of variable with values (runtime only)
type Terminal struct {
	Variables []Variable
}

// Variable - container holding variable values
type Variable struct {
	VariableName string      `json:"name"`
	VariableData interface{} `json:"data"`
	VariableType string      `json:"type"`
}

// NewTerminal - attempts to start io handler for term commands
func NewTerminal() error {
	reader := bufio.NewScanner(os.Stdin) // Init reader

	nodeClient := nodeProto.NewNodeProtobufClient("http://localhost:8080", &http.Client{})          // Init node client
	handlerClient := handlerProto.NewHandlerProtobufClient("http://localhost:8080", &http.Client{}) // Init handler client

	for {
		fmt.Print("\n> ") // Print prompt

		reader.Scan() // Scan

		input := reader.Text() // Fetch string input

		input = strings.TrimSuffix(input, "\n") // Trim newline

		receiver, methodname, params, err := common.ParseStringMethodCall(input) // Attempt to parse as method call

		if err != nil { // Check for errors
			panic(err) // Panic
		}

		switch receiver {
		case "node":
			err := handleNode(&nodeClient, methodname, params) // Handle node

			if err != nil { // Check for errors
				fmt.Println(err.Error()) // Log found error
			}
		case "handler":
			err := handleHandler(&handlerClient, methodname, params) // Handle handler

			if err != nil { // Check for errors
				fmt.Println(err.Error()) // Log found error
			}
		}
	}
}

func handleNode(nodeClient *nodeProto.Node, methodname string, params []string) error {
	if len(params) == 0 { // Check for nil parameters
		return errors.New("invalid parameters (requires at least 1 parameter)") // Return error
	}

	reflectParams := []reflect.Value{} // Init buffer

	reflectParams = append(reflectParams, reflect.ValueOf(context.Background())) // Append request context

	switch methodname {
	case "NewNode":
		if len(params) != 2 { // Check for insufficient parameters
			return errors.New("invalid parameters (requires string, int)") // Return error
		}

		boolVal, _ := strconv.ParseBool(params[1]) // Parse isBootstrap

		reflectParams = append(reflectParams, reflect.ValueOf(&nodeProto.GeneralRequest{Address: params[0], IsBootstrap: boolVal})) // Append params
	case "StartListener":
		intVal, _ := strconv.Atoi(params[0]) // Get int val

		reflectParams = append(reflectParams, reflect.ValueOf(&nodeProto.GeneralRequest{Port: uint32(intVal)})) // Append params
	case "WriteToMemory", "ReadFromMemory":
		reflectParams = append(reflectParams, reflect.ValueOf(&nodeProto.GeneralRequest{Path: params[0]})) // Append params
	default:
		return errors.New("illegal method " + methodname)
	}

	result := reflect.ValueOf(*nodeClient).MethodByName(methodname).Call(reflectParams) // Call method

	response := result[0].Interface().(*nodeProto.GeneralResponse) // Get response

	if result[1].Interface() != nil { // Check for errors
		fmt.Println(result[1].Interface().(error).Error()) // Log error
	} else {
		fmt.Println(response.Message) // Log response
	}

	return nil // No error occurred, return nil
}

func handleHandler(handlerClient *handlerProto.Handler, methodname string, params []string) error {
	if len(params) == 0 { // Check for nil parameters
		return errors.New("invalid parameters") // Return error
	}

	reflectParams := []reflect.Value{} // Init buffer

	reflectParams = append(reflectParams, reflect.ValueOf(context.Background())) // Append request context

	switch methodname {
	case "StartHandler":
		port, _ := strconv.Atoi(params[0]) // Parse port

		reflectParams = append(reflectParams, reflect.ValueOf(&handlerProto.GeneralRequest{Port: uint32(port)})) // Append params
	default:
		return errors.New("illegal method " + methodname)
	}

	result := reflect.ValueOf(*handlerClient).MethodByName(methodname).Call(reflectParams) // Call method

	response := result[0].Interface().(*handlerProto.GeneralResponse) // Get response

	if result[1].Interface() != nil { // Check for errors
		fmt.Println(result[1].Interface().(error).Error()) // Log error
	} else {
		fmt.Println(response.Message) // Log response
	}

	return nil // No error occurred, return nil
}

// AddVariable - attempt to append specified variable to terminal variable list
func (term *Terminal) AddVariable(variableName string, variableData interface{}, variableType string) error {
	variable := Variable{VariableName: variableName, VariableData: variableData, VariableType: variableType}

	if reflect.ValueOf(term).IsNil() { // Check for nil variable
		return errors.New("nil terminal found") // Return error
	}

	if len(term.Variables) == 0 { // Check for uninitialized variable array
		term.Variables = []Variable{variable} // Initialize with variable

		return nil // No error occurred, return nil
	}

	term.Variables = append(term.Variables, variable) // Append to array

	return nil // No error occurred, return nil
}

// ReplaceVariable - attempt to replace value at index with specified variable
func (term *Terminal) ReplaceVariable(variableIndex int, variableData interface{}) error {
	if reflect.ValueOf(term).IsNil() { // Check for nil variable
		return errors.New("nil terminal found") // Return error
	}

	if len(term.Variables) == 0 { // Check for uninitialized variable array
		return errors.New("empty terminal environment") // Return found error
	}

	(*term).Variables[variableIndex].VariableData = variableData // Replace value

	return nil
}

// QueryType - attempt to fetch index of variable with matching type
func (term *Terminal) QueryType(variableType string) (uint, error) {
	if variableType == "" { // Check for nil parameter
		return 0, errors.New("invalid type") // Return found error
	}

	if len(term.Variables) == 0 { // Check that terminal environment is not nil
		return 0, errors.New("empty terminal environment") // Return found error
	}

	for x := 0; x != len(term.Variables); x++ { // Declare, increment iterator
		if term.Variables[x].VariableType == variableType { // Check for match
			return uint(x), nil // Return result
		}
	}

	return 0, errors.New("couldn't find matching variable") // Return error
}

// hasVariableSet - checks if specified command sets a variable
func hasVariableSet(command string) bool {
	if strings.HasPrefix(strings.ToLower(command), "var") { // Check for prefix
		return true
	}

	return false
}
