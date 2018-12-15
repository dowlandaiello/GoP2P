package connection

import (
	"strings"
	"testing"

	"github.com/dowlandaiello/GoP2P/common"
)

func TestFromBytes(t *testing.T) {
	connection, err := generateConnection() // Create connection

	if err != nil && !strings.Contains(err.Error(), "socket") { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	} else if err != nil && strings.Contains(err.Error(), "socket") {
		t.Logf("WARNING: socket actions require sudo privileges.") // Log warning
	}

	serializedConnection, err := common.SerializeToBytes(*connection) // Serialize connection

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
