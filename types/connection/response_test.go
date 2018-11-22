package connection

import (
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
)

func TestResponseFromBytes(t *testing.T) {
	response := Response{[][]byte{[]byte("test")}} // Create instance of response{} struct

	serializedResponse, err := common.SerializeToBytes(response) // Attempt to serialize response to byte array

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	readResponse, err := ResponseFromBytes(serializedResponse) // Attempt to read bytes

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("found response %s", readResponse) // Log success
}
