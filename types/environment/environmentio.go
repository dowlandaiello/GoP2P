package environment

import (
	"path/filepath"

	"github.com/mitsukomegumi/GoP2P/common"
)

// WriteToMemory - create serialized instance of specified environment in specified path (string)
func (environment *Environment) WriteToMemory(path string) error {
	err := common.WriteGob(path+filepath.FromSlash("/environment.gob"), environment) // Attempt to write env to path

	if err != nil { // Check for errors
		return err // Return error
	}

	return nil // No error occurred, return nil.
}

// ReadEnvironmentFromMemory - read serialized object of specified node database from specified path
func ReadEnvironmentFromMemory(path string) (*Environment, error) {
	tempEnv := new(Environment)

	err := common.ReadGob(path+filepath.FromSlash("/environment.gob"), tempEnv)
	if err != nil { // Check for errors
		return nil, err // Return error
	}
	return tempEnv, nil // No error occurred, return nil error, env
}
