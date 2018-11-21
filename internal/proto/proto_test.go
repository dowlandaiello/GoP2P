package proto

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
)

// TestNewProtobufMessage - test functionality of NewProtobufMessage() method
func TestNewProtobufMessage(t *testing.T) {
	_, err := NewProtobufGuide("test.proto", "test") // Init guide

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	test := &Test{ // Init test struct
		Test:  "test",
		Test2: "test2",
	}

	data, err := proto.Marshal(test) // Marshal

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	protoMessage, err := NewProtobufMessage("test.proto", data) // Init message

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	marshaledVal, err := json.Marshal(*protoMessage) // Marshal message

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf(string(marshaledVal)) // Log success
}
