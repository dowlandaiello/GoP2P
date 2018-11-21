package proto

import (
	"encoding/json"
	"testing"
)

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
