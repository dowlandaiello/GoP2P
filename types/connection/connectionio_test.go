package connection

import (
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
)

func TestFromBytes(t *testing.T) {
	connection, err := generateConnection() // Create connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	serializedConnection, err := common.SerializeToBytes(connection) // Serialize connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	readConnection, err := FromBytes(serializedConnection) // Attempt to read from serialized connection

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("Decoded connection with address %s", readConnection.InitializationNode.Address) // Log success
}
