package proto

import (
	"github.com/dowlandaiello/GoP2P/common"
	"github.com/dowlandaiello/GoP2P/types/shard"
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

// SendToShard - shard.SendBytesShard() wrapper
func (protoMessage *ProtobufMessage) SendToShard(shardAddress string, port int) error {
	serialized, err := protoMessage.ToBytes() // Serialize to bytes

	if err != nil { // Check for errors
		return err // Return found error
	}

	return shard.SendBytesShard(serialized, shardAddress, port) // Send to address
}

// SendToShardResult - shard.SendBytesShardResult() wrapper
func (protoMessage *ProtobufMessage) SendToShardResult(shardAddress string, port int) ([]byte, error) {
	serialized, err := protoMessage.ToBytes() // Serialize to bytes

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	return shard.SendBytesShardResult(serialized, shardAddress, port) // Send to address
}
