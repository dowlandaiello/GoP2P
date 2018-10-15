package database

import (
	"bytes"
	"encoding/json"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
)

// WriteToMemory - create serialized instance of specified NodeDatabase in specified path (string)
func (db *NodeDatabase) WriteToMemory(env *environment.Environment) error {
	variable, err := environment.NewVariable(db.NetworkAlias+"NodeDatabase", *db)

	if err != nil { // Check for errors
		return err // Return error
	}

	err = env.AddVariable(variable, false) // Attempt to add specified variable

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil.
}

// ReadDatabaseFromMemory - read serialized object of specified node database from specified path
func ReadDatabaseFromMemory(env *environment.Environment, networkAlias string) (*NodeDatabase, error) {
	variable, err := env.QueryType(networkAlias + "NodeDatabase") // Attempt to fetch db

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	database := NodeDatabase{} // Init buffer

	decoded, err := common.InterfaceFromBytes(variable.VariableData, &database) // Fetch value

	if err != nil { // Check for errors
		return &NodeDatabase{}, err // Return found error
	}

	db := decoded.(*NodeDatabase)

	return db, nil // No error occurred, return nil error, db
}

// FromBytes - attempt to convert specified byte array to db
func FromBytes(b []byte) (*NodeDatabase, error) {
	object := NodeDatabase{} // Create empty instance

	err := json.NewDecoder(bytes.NewReader(b)).Decode(&object) // Attempt to read

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return &object, nil // No error occurred, return read value
}
