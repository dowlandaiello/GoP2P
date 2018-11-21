package proto

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// ProtobufMessage - container holding protobuf message, metadata/supporting proto file
type ProtobufMessage struct {
	Message []byte // Message - serialized protobuf message

	Guide *ProtobufGuide // Guide - messag protobuf guide
}

// ProtobufGuide - container holding metadata/supporting proto file
type ProtobufGuide struct {
	ProtoID string // ProtoID - determines 'class' of protofile (specified by user)

	ProtoGuide []byte // ProtoGuide - serialized proto file, used for reading of protobuf message
	GoGuide    []byte // GoGuide - serialized golang support file, used for reading of protobuf message
}

/* BEGIN EXPORTED METHODS */

// NewProtobufMessage - initialize protobuf message
func NewProtobufMessage(guidePath string, data []byte) (*ProtobufMessage, error) {
	protoGuide, err := ReadGuideFromMemory(guidePath) // Read guide

	if err != nil { // Check for errors
		return &ProtobufMessage{}, err // Return found error
	}

	protoMessage := &ProtobufMessage{ // Init message
		Message: data,
		Guide:   protoGuide,
	}

	return protoMessage, nil // No error occurred, return nil
}

// NewProtobufGuide - initialize and register protobuf message guide
func NewProtobufGuide(protofilePath string, protoID string) (*ProtobufGuide, error) {
	readProtofile, err := ioutil.ReadFile(protofilePath) // Read protofile

	if err != nil { // Check for errors
		return &ProtobufGuide{}, err // Return found error
	}

	protoGuide := &ProtobufGuide{ // Init guide
		ProtoID:    protoID,
		ProtoGuide: readProtofile,
	}

	err = protoGuide.register(protofilePath, protoID) // Register guide

	if err != nil { // Check for errors
		return &ProtobufGuide{}, err //  Return found error
	}

	return protoGuide, nil // No error occurred, return nil
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

// register - write protobufGuide to memory for later use, generate necessary go support files
func (protoGuide *ProtobufGuide) register(protofilePath string, protoID string) error {
	err := protoGuide.WriteToMemory(fmt.Sprintf("%s.goP2PGuide", protofilePath)) // Write to memory at given path

	if err != nil { // Check for errors
		return err // Return found error
	}

	err = generateProto(protofilePath) // Generate proto

	if err != nil { // Check for errors
		return err // Return found error
	}

	supportPath := strings.Split(protofilePath, ".proto")[0] // Trim .proto suffix

	readSupportfile, err := ioutil.ReadFile(fmt.Sprintf("%s.pb.go", supportPath)) // Read support file

	if err != nil { // Check for errors
		return err // Return found error
	}

	(*protoGuide).GoGuide = readSupportfile // Set support file

	return nil // No error occurred, return nil
}

// generateProto - read specified protofile, generate necessary .pb.go files (protoc required)
func generateProto(protofilePath string) error {
	args := []string{ // Init arguments
		"--go_out=.",
		protofilePath,
	}

	cmd := exec.Command("protoc", args...) // Init exec command

	return cmd.Run() // Return cmd.Run() output
}

/* END INTERNAL METHODS */
