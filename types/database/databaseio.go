package database

import (
	"path/filepath"

	"github.com/mitsukomegumi/GoP2P/common"
)

// WriteToMemory - create serialized instance of specified NodeDatabase in specified path (string)
func (db *NodeDatabase) WriteToMemory(path string) error {
	err := common.WriteGob(path+filepath.FromSlash("/nodeDb.gob"), db) // Attempt to write db to path

	if err != nil { // Check for errors
		return err // Return error
	}

	return nil // No error occurred, return nil.
}

// ReadDatabaseFromMemory - read serialized object of specified node database from specified path
func ReadDatabaseFromMemory(path string) (*NodeDatabase, error) {
	tempDb := new(NodeDatabase)

	err := common.ReadGob(path+filepath.FromSlash("/nodeDb.gob"), tempDb)
	if err != nil { // Check for errors
		return nil, err // Return error
	}
	return tempDb, nil // No error occurred, return nil error, db
}
