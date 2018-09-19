package common

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"strings"

	"golang.org/x/crypto/sha3"
)

const (
	// RootCallEndpoint - global const defining the base endpoint for calls
	RootCallEndpoint = "/call"
)

/* BEGIN EXPORTED METHODS */

// HelloWorld - a little hello go-rpcify call
func HelloWorld() (string, error) {
	return "Hello, world!", nil
}

/*
	BEGIN CRYPTO METHODS:
*/

// SHA3String - generate and return SHA3 string-value hash of specified byte array
func SHA3String(b []byte) string {
	hash := sha3.Sum256(b) // Calculate hash

	return base64.StdEncoding.EncodeToString(hash[:]) // Return base64 string
}

// SHA3URLSafeString - generate and return SHA3 string-value hash of specified byte array
func SHA3URLSafeString(b []byte) string {
	hash := sha3.Sum256(b) // Calculate hash

	return strings.Split(base64.StdEncoding.EncodeToString(hash[:]), "/")[0] // Return base64 string
}

// SHA3Bytes - generate and return SHA3 hash of specified byte array
func SHA3Bytes(b []byte) []byte {
	hash := sha3.Sum256(b) // Calculate hash

	return hash[:] // Return hash
}

/*
	END CRYPTO METHODS
*/

/*
	BEGIN CONVERSION METHODS:
*/

// ToBytes - attempt to encode specified interface to byte array
func ToBytes(object interface{}) ([]byte, error) {
	var buf bytes.Buffer // Init buffer

	encoder := gob.NewEncoder(&buf) // Init encoder

	err := encoder.Encode(object) // Attempt to encode

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return buf.Bytes(), nil // No error occurred, return byte value
}

/*
	END CONVERSION METHODS
*/

/* END EXPORTED METHODS */
