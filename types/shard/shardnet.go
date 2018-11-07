package shard

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mitsukomegumi/GoP2P/common"
)

// SendBytesShardResult - attempt to send specified bytes to given shard address, returning result
func SendBytesShardResult(b []byte, address string, port int) ([]byte, error) {
	if len(address) == 0 || len(address) < 0 || !strings.Contains(address, "::") { // Check for invalid input
		return []byte{}, fmt.Errorf("invalid address %s", address) // Return found error
	}

	addresses := strings.Split(strings.Split(address, "::")[1], ":") // Split into string slice

	buffer := [][]byte{} // Init buffer

	finished := make(chan []bool) // Init finished

	startTime := time.Now() // Get current time

	for _, address := range addresses { // Iterate through addresses
		if float64(len(finished)) > (0.51 * float64(len(addresses))) { // Check 51% of nodes finished
			filteredResult, err := common.GetCommonByteDifference(buffer) // Fetch final result

			if err != nil { // Check for errors
				return []byte{}, err // Return found error
			}

			return filteredResult, nil // Return read data
		}

		address = address + ":" + strconv.Itoa(port) // Append port

		go common.SendBytesResultBufferAsync(b, buffer, address, finished) // Send to address, append to buffer
	}

	for float64(len(finished)) < (0.51 * float64(len(addresses))) { // Check 51% of nodes finished
		if time.Now().Sub(startTime) > 3*time.Second { // Check for time out
			return []byte{}, errors.New("timed out") // Return timed out
		}
	}

	filteredResult, err := common.GetCommonByteDifference(buffer) // Fetch final result

	if err != nil { // Check for errors
		return []byte{}, err // Return found error
	}

	return filteredResult, nil // Return read data
}

// SendBytesShard - attempt to send specified bytes to given shard address
func SendBytesShard(b []byte, address string, port int) error {
	if len(address) == 0 || len(address) < 0 || !strings.Contains(address, "::") { // Check for invalid input
		return fmt.Errorf("invalid address %s", address) // Return found error
	}

	addresses := strings.Split(strings.Split(address, "::")[1], ":") // Split into string slice

	finished := []bool{} // Init finished buffer

	startTime := time.Now() // Fetch current time

	for _, address := range addresses { // Iterate through addresses
		address = address + ":" + strconv.Itoa(port) // Append port

		go common.SendBytesAsync(b, address, finished) // Send to address

		if len(finished) >= len(addresses) { // Check finished
			return nil // No error occurred, return nil
		}
	}

	for float64(len(finished)) < 0.75*float64(len(addresses)) { // Check wrote to all nodes
		if time.Now().Sub(startTime) > 10*time.Second { // Check for timeout
			return errors.New("timed out") // Return timed out
		}
	}

	return nil // No error occurred, return  nil
}
