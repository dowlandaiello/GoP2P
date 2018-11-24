package proto

import (
	"strconv"
	"strings"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/shard"
)

// SendToAddress - common.SendBytes() wrapper
func (protoMessage *ProtobufMessage) SendToAddress(address string) error {
	serialized, err := protoMessage.ToBytes() // Serialize to bytes

	if err != nil { // Check for errors
		return err // Return found error
	}

	return common.SendBytes(serialized, address) // Send to address
}

// SendToAddressResult - common.SendBytesResult() wrapper
func (protoMessage *ProtobufMessage) SendToAddressResult(address string) ([]byte, error) {
	serialized, err := protoMessage.ToBytes() // Serialize to bytes

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	return common.SendBytesResult(serialized, address) // Send to address
}

// SendToShardResult - shard.SendBytesShardResult() wrapper
func (protoMessage *ProtobufMessage) SendToShardResult(shardAddress string) ([]byte, error) {
	serialized, err := protoMessage.ToBytes() // Serialize to bytes

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	port, err := strconv.Atoi(strings.Split(shardAddress, ":")[1]) // Get port

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	return shard.SendBytesShardResult(serialized, shardAddress, port) // Send to address
}

// SendToShard - shard.SendBytesShard() wrapper
func (protoMessage *ProtobufMessage) SendToShard(shardAddress string) error {
	serialized, err := protoMessage.ToBytes() // Serialize to bytes

	if err != nil { // Check for errors
		return err // Return found error
	}

	port, err := strconv.Atoi(strings.Split(shardAddress, ":")[1]) // Get port

	if err != nil { // Check for errors
		return err // Return found error
	}

	return shard.SendBytesShard(serialized, shardAddress, port) // Send to address
}
