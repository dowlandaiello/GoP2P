package proto

import (
	"encoding/json"
	"testing"

	proto "github.com/golang/protobuf/proto"
)

// TestToBytes - test functionality of ToBytes() method
func TestToBytes(t *testing.T) {
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

	data, err = protoMessage.ToBytes() // Marshal to bytes

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf(string(data)) // Log success
}

// TestFromBytes - test functionality of FromBytes() method
func TestFromBytes(t *testing.T) {
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

	data, err = protoMessage.ToBytes() // Marshal to bytes

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	message, err := FromBytes(data) // Decode message

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf(message.String()) // Log success
}

// TestReadGuideFromMemory - test functionality of ReadGuideFromMemory() method
func TestReadGuideFromMemory(t *testing.T) {
	protoGuide, err := NewProtobufGuide("test.proto", "test") // Init guide

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = protoGuide.WriteToMemory("test.proto.goP2PGuide") // Write guide to memory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	guide, err := ReadGuideFromMemory("test.proto") // Read guide from memory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	marshaledVal, err := json.MarshalIndent(*guide, "", "  ") // Marshal guide

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf(string(marshaledVal)) // Log success
}

// TestWriteGuideToMemory - test functionality of WriteGuideToMemory() function
func TestWriteGuideToMemory(t *testing.T) {
	protoGuide, err := NewProtobufGuide("test.proto", "test") // Init guide

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	err = protoGuide.WriteToMemory("test.proto.goP2PGuide") // Write guide to memory

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	marshaledVal, err := json.MarshalIndent(*protoGuide, "", "  ") // Marshal guide

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf(string(marshaledVal)) // Log success
}
