package proto

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
)

// TestSendToAddress - test functionality of SendToAddress() method
func TestSendToAddress(t *testing.T) {
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

	err = protoMessage.SendToAddress("1.1.1.1:443") // Send message

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}

// TestSendToAddressResult - test functionality of TestSendToAddressResult() method
func TestSendToAddressResult(t *testing.T) {
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

	result, err := protoMessage.SendToAddressResult("1.1.1.1:443") // Send message

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf(string(result)) // Log success
}

// TestSendToShard - test functionality of SendToShard() method
func TestSendToShard(t *testing.T) {
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

	err = protoMessage.SendToShard("1.1.1.1::1.1.1.1", 443) // Send message

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}
}

// TestSendToShardResult - test functionality of SendToShardResult() method
func TestSendToShardResult(t *testing.T) {
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

	result, err := protoMessage.SendToShardResult("1.1.1.1::1.1.1.1", 443) // Send message

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf(string(result)) // Log success
}
