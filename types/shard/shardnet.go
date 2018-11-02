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
func SendBytesShardResult(b []byte, address string, port int) ([][]byte, error) {
	if len(address) == 0 || len(address) < 0 || !strings.Contains(address, "::") {
		return [][]byte{}, fmt.Errorf("invalid address %s", address) // Return found error
	}

	addresses := strings.Split(strings.Split(address, "::")[1], ":") // Split into string slice

	buffer := [][]byte{} // Init buffer

	finished := make(chan []bool) // Init finished

	startTime := time.Now() // Get current time

	for _, address := range addresses { // Iterate through addresses
		if float64(len(finished)) > (0.51 * float64(len(addresses))) { // Check 51% of nodes finished
			return buffer, nil // Return read data
		}

		address = address + ":" + strconv.Itoa(port) // Append port

		go common.SendBytesResultBufferAsync(b, buffer, address, finished) // Send to address, append to buffer
	}

	for float64(len(finished)) < (0.51 * float64(len(addresses))) {
		fmt.Println(len(finished))
		fmt.Println(0.51 * float64(len(addresses)))

		if time.Now().Sub(startTime) > 3*time.Second { // Check for time out
			return [][]byte{}, errors.New("timed out") // Return timed out
		}
	}

	return buffer, nil // Return read data
}
