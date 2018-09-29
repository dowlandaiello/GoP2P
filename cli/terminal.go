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
	commonProto "github.com/mitsukomegumi/GoP2P/rpc/proto/common"
	databaseProto "github.com/mitsukomegumi/GoP2P/rpc/proto/database"
	environmentProto "github.com/mitsukomegumi/GoP2P/rpc/proto/environment"
	handlerProto "github.com/mitsukomegumi/GoP2P/rpc/proto/handler"
	nodeProto "github.com/mitsukomegumi/GoP2P/rpc/proto/node"
	upnpProto "github.com/mitsukomegumi/GoP2P/rpc/proto/upnp"
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

	nodeClient := nodeProto.NewNodeProtobufClient("http://localhost:8080", &http.Client{})                      // Init node client
	handlerClient := handlerProto.NewHandlerProtobufClient("http://localhost:8080", &http.Client{})             // Init handler client
	environmentClient := environmentProto.NewEnvironmentProtobufClient("http://localhost:8080", &http.Client{}) // Init environment client
	upnpClient := upnpProto.NewUpnpProtobufClient("http://localhost:8080", &http.Client{})                      // Init upnp client
	databaseClient := databaseProto.NewDatabaseProtobufClient("http://localhost:8080", &http.Client{})          // Init database client
	commonClient := commonProto.NewCommonProtobufClient("http://localhost:8080", &http.Client{})                // Init common client

	for {
		fmt.Print("\n> ") // Print prompt

		reader.Scan() // Scan

		input := reader.Text() // Fetch string input

		input = strings.TrimSuffix(input, "\n") // Trim newline

		receiver, methodname, params, err := common.ParseStringMethodCall(input) // Attempt to parse as method call

		if err != nil { // Check for errors
			fmt.Println(err.Error()) // Log found error

			continue
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
		case "environment":
			err := handleEnvironment(&environmentClient, methodname, params) // Handle environment

			if err != nil { // Check for errors
				fmt.Println(err.Error()) // Log found error
			}
		case "upnp":
			err := handleUpnp(&upnpClient, methodname, params) // Handle upnp

			if err != nil { // Check for errors
				fmt.Println(err.Error()) // Log found error
			}
		case "database":
			err := handleDatabase(&databaseClient, methodname, params) // Handle database

			if err != nil { // Check for errors
				fmt.Println(err.Error()) // Log found error
			}
		case "common":
			err := handleCommon(&commonClient, methodname, params) // Handle common

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
		return errors.New("illegal method: " + methodname + ", available methods: NewNode(), StartListener(), WriteToMemory(), ReadFromMemory()") // Return error
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
		return errors.New("invalid parameters (requires at least 1 parameter)") // Return error
	}

	reflectParams := []reflect.Value{} // Init buffer

	reflectParams = append(reflectParams, reflect.ValueOf(context.Background())) // Append request context

	switch methodname {
	case "StartHandler":
		port, _ := strconv.Atoi(params[0]) // Parse port

		reflectParams = append(reflectParams, reflect.ValueOf(&handlerProto.GeneralRequest{Port: uint32(port)})) // Append params
	default:
		return errors.New("illegal method: " + methodname + ", available methods: StartHandler()") // Return error
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

func handleEnvironment(environmentClient *environmentProto.Environment, methodname string, params []string) error {
	reflectParams := []reflect.Value{} // Init buffer

	reflectParams = append(reflectParams, reflect.ValueOf(context.Background())) // Append request context

	switch methodname {
	case "NewEnvironment":
		reflectParams = append(reflectParams, reflect.ValueOf(&environmentProto.GeneralRequest{})) // Append empty request
	case "QueryType":
		if len(params) != 1 { // Check for errors
			return errors.New("invalid parameters (requires string)") // Return found error
		}

		queryTypeVal := params[0] // Fetch queryTypeVal

		reflectParams = append(reflectParams, reflect.ValueOf(&environmentProto.GeneralRequest{VariableType: queryTypeVal})) // Append querytype request
	case "QueryValue":
		if len(params) != 1 { // Check for errors
			return errors.New("invalid parameters (requires string)") // Return found error
		}

		queryValueVal := params[0] // Fetch query val

		reflectParams = append(reflectParams, reflect.ValueOf(&environmentProto.GeneralRequest{Value: queryValueVal})) // Append queryval request
	case "NewVariable":
		if len(params) != 2 { // Check for errors
			return errors.New("invalid parameters (requires string, string)") // Return found error
		}

		variablePathVal := params[0] // Fetch variable data path
		variableTypeVal := params[1] // Fetch variable type

		reflectParams = append(reflectParams, reflect.ValueOf(&environmentProto.GeneralRequest{Path: variablePathVal, VariableType: variableTypeVal})) // Append path request
	case "AddVariable", "WriteToMemory", "ReadFromMemory":
		if len(params) != 1 { // Check for errors
			return errors.New("invalid parameters (requires string)") // Return found error
		}

		pathVal := params[0] // Fetch variable data path

		reflectParams = append(reflectParams, reflect.ValueOf(&environmentProto.GeneralRequest{Path: pathVal})) // Append path request
	default:
		return errors.New("illegal method: " + methodname + ", available methods: NewEnvironment(), QueryType(), QueryValue(), NewVariable(), AddVariable(), WriteToMemory(), ReadFromMemory()") // Return error
	}

	result := reflect.ValueOf(*environmentClient).MethodByName(methodname).Call(reflectParams) // Call method

	response := result[0].Interface().(*environmentProto.GeneralResponse) // Get response

	if result[1].Interface() != nil { // Check for errors
		fmt.Println(result[1].Interface().(error).Error()) // Log error
	} else {
		fmt.Println(response.Message) // Log response
	}

	return nil // No error occurred, return nil
}

func handleUpnp(upnpClient *upnpProto.Upnp, methodname string, params []string) error {
	reflectParams := []reflect.Value{} // Init buffer

	reflectParams = append(reflectParams, reflect.ValueOf(context.Background())) // Append request context

	switch methodname {
	case "GetGateway":
		reflectParams = append(reflectParams, reflect.ValueOf(&upnpProto.GeneralRequest{})) // Append params
	case "ForwardPortSilent", "ForwardPort", "RemoveForwarding":
		if len(params) != 1 { // Check for invalid parameters
			return errors.New("invalid parameters (requires uint32)") // Return error
		}

		port, err := strconv.Atoi(params[0]) // Convert to int

		if err != nil { // Check for errors
			return err // Return found error
		}

		reflectParams = append(reflectParams, reflect.ValueOf(&upnpProto.GeneralRequest{PortNumber: uint32(port)})) // Append params
	default:
		return errors.New("illegal method: " + methodname + ", available methods: GetGateway(), ForwardPortSilent(), ForwardPort(), RemoveForwarding()") // Return error
	}

	result := reflect.ValueOf(*upnpClient).MethodByName(methodname).Call(reflectParams) // Call method

	response := result[0].Interface().(*upnpProto.GeneralResponse) // Get response

	if result[1].Interface() != nil { // Check for errors
		fmt.Println(result[1].Interface().(error).Error()) // Log error
	} else {
		fmt.Println(response.Message) // Log response
	}

	return nil // No error occurred, return nil
}

func handleDatabase(databaseClient *databaseProto.Database, methodname string, params []string) error {
	reflectParams := []reflect.Value{} // Init buffer

	reflectParams = append(reflectParams, reflect.ValueOf(context.Background())) // Append request context

	switch methodname {
	case "NewDatabase":
		if len(params) != 1 { // Check for invalid parameters
			return errors.New("invalid parameters (requires uint32)") // Return error
		}

		acceptableTimeout, err := strconv.Atoi(params[0]) // Fetch acceptable timeout

		if err != nil { // Check for errors
			return err // Return found error
		}

		reflectParams = append(reflectParams, reflect.ValueOf(&databaseProto.GeneralRequest{AcceptableTimeout: uint32(acceptableTimeout)})) // Append params
	case "AddNode":
		reflectParams = append(reflectParams, reflect.ValueOf(&databaseProto.GeneralRequest{})) // Append nil params
	case "RemoveNode", "QueryForAddress":
		if len(params) != 1 { // Check for invalid parameters
			return errors.New("invalid parameters (requires string)") // Return error
		}

		address := params[0] // Fetch removal address

		reflectParams = append(reflectParams, reflect.ValueOf(&databaseProto.GeneralRequest{Address: address})) // Append params
	case "WriteToMemory", "ReadFromMemory":
		if len(params) != 1 { // Check for invalid parameters
			return errors.New("invalid parameters (requires string)") // Return error
		}

		path := params[0] // Fetch path

		reflectParams = append(reflectParams, reflect.ValueOf(&databaseProto.GeneralRequest{DataPath: path})) // Append params
	case "FromBytes":
		if len(params) != 1 { // Check for invalid parameters
			return errors.New("invalid parameters (requires []byte)") // Return error
		}

		byteVal := []byte(params[0]) // Fetch byte val

		reflectParams = append(reflectParams, reflect.ValueOf(&databaseProto.GeneralRequest{ByteVal: byteVal})) // Append params
	default:
		return errors.New("illegal method: " + methodname + ", available methods: NewDatabase(), AddNode(), RemoveNode(), QueryForAddress(), WriteToMemory(), ReadFromMemory(), FromBytes()") // Return error
	}

	result := reflect.ValueOf(*databaseClient).MethodByName(methodname).Call(reflectParams) // Call method

	response := result[0].Interface().(*databaseProto.GeneralResponse) // Get response

	if result[1].Interface() != nil { // Check for errors
		fmt.Println(result[1].Interface().(error).Error()) // Log error
	} else {
		fmt.Println(response.Message) // Log response
	}

	return nil // No error occurred, return nil
}

func handleCommon(commonClient *commonProto.Common, methodname string, params []string) error {
	reflectParams := []reflect.Value{} // Init buffer

	reflectParams = append(reflectParams, reflect.ValueOf(context.Background())) // Append request context

	switch methodname {
	case "ParseStringMethodCall", "ParseStringParams", "StringStripReceiverCall", "StringStripParentheses", "StringFetchCallReceiver", "CheckAddress":
		if len(params) != 1 { // Check for invalid parameters
			return errors.New("invalid parameters (requires string)") // Return error
		}

		reflectParams = append(reflectParams, reflect.ValueOf(&commonProto.GeneralRequest{Input: params[0]})) // Append params
	case "ConvertStringToReflectValues":
		if len(params) < 1 { // Check for invalid parameters
			return errors.New("invalid parameters (requires string)") // Return error
		}

		reflectParams = append(reflectParams, reflect.ValueOf(&commonProto.GeneralRequest{Inputs: params})) // Append params
	case "SHA256":
		if len(params) != 1 { // Check for invalid parameters
			return errors.New("invalid parameters (requires string)")
		}

		reflectParams = append(reflectParams, reflect.ValueOf(&commonProto.GeneralRequest{ByteInput: []byte(params[0])})) // Append params
	case "SendBytes":
		if len(params) != 2 { // Check for invalid parameters
			return errors.New("invalid parameters (requires []byte, string)")
		}

		reflectParams = append(reflectParams, reflect.ValueOf(&commonProto.GeneralRequest{ByteInput: []byte(params[0]), Input: params[1]})) // Append params
	case "GetExtIPAddrWithUpNP", "GetExtIPAddrWithoutUpNP", "GetCurrentTime", "GetCurrentDir":
		reflectParams = append(reflectParams, reflect.ValueOf(&commonProto.GeneralRequest{})) // Append empty params
	default:
		return errors.New("illegal method: " + methodname + ", available methods: ParseStringMethodCall(), ParseStringParams(), StringStripReceiverCall(), StringStripParentheses(), StringFetchCallReceiver(), CheckAddress(), ConvertStringToReflectValues(), SHA256(), SendBytes(), GetExtIPAddrWithUpnp(), GetExtIPAddrWithoutUpnp(), GetCurrentTime(), GetCurrentDir()") // Return error
	}

	result := reflect.ValueOf(*commonClient).MethodByName(methodname).Call(reflectParams) // Call method

	response := result[0].Interface().(*commonProto.GeneralResponse) // Get response

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
