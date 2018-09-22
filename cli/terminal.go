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
	proto "github.com/mitsukomegumi/GoP2P/rpc/proto"
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
	reader := bufio.NewReader(os.Stdin) // Init reader

	nodeClient := proto.NewNodeProtobufClient("http://localhost:8080", &http.Client{}) // Init node client

	for {
		fmt.Print("\n> ")

		input, err := reader.ReadString('\n') // Search for user input

		if err != nil { // Check for errors
			panic(err) // Panic
		}

		receiver, methodname, params, err := common.ParseStringMethodCall(input) // Attempt to parse as method call

		if err != nil { // Check for errors
			panic(err) // Panic
		}

		switch receiver {
		case "node":
			reflectParams := []reflect.Value{} // Init buffer

			switch methodname {
			case "NewNode":
				boolVal, _ := strconv.ParseBool(params[1]) // Parse isBootstrap

				reflectParams = append(reflectParams, reflect.ValueOf(context.Background())) // Append request context

				reflectParams = append(reflectParams, reflect.ValueOf(&proto.GeneralRequest{Address: params[0], IsBootstrap: boolVal})) // Append params
			}

			result := reflect.ValueOf(nodeClient).MethodByName(methodname).Call(reflectParams) // Call method

			response := result[0].Interface().(*proto.GeneralResponse) // Get response

			if result[1].Interface() != nil { // Check for errors
				fmt.Println(result[1].Interface().(error).Error()) // Log error
			} else {
				fmt.Println(response.Message) // Log response
			}
		}
	}
}

func handleNode(nodeClient *proto.Node) {

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
