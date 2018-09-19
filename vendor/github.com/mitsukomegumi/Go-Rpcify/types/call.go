package types

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mitsukomegumi/Go-Rpcify/common"
)

// Call - call to specified method
type Call struct {
	Method     func() (string, error) `json:"method"` // Used to call method
	MethodHash string                 `json:"hash"`   // Used to identify call

	Endpoint string `json:"endpoint"` // Used for calls to rpc
}

/* BEGIN EXPORTED METHODS */

// NewCall - initialize new instance of Call struct
func NewCall(method func() (string, error), endpoint string) (*Call, error) {
	if reflect.ValueOf(method).IsNil() { // Check for nil method
		return &Call{}, errors.New("nil call") // Return error
	}

	methodHash := ""

	call := Call{method, methodHash, endpoint} // Init call

	byteValue, err := common.ToBytes(fmt.Sprintf("%v", call)) // Attempt to encode

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	call.MethodHash = common.SHA3URLSafeString(byteValue) // Calculate method hash

	if call.Endpoint == "" { // Check for nil endpoint
		call.Endpoint = common.RootCallEndpoint + "/" + call.MethodHash // Set endpoint
	}

	return &call, nil // Return initialized call instance
}

// Run - attempt to run specified call
func (call *Call) Run() (string, error) {
	output, err := call.Method() // Run method

	if err != nil { // Check for errors
		return "", err // Return found error
	}

	return output, nil // No error occurred, return nil
}

/* END EXPORTED METHODS */
