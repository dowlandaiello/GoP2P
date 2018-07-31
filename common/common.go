package common

import (
	"net"
	"time"

	"github.com/mitsukomegumi/GoP2P/fastping"
)

const (
	// NodeAvailableRep - global definition for value of node availability
	NodeAvailableRep = 10
)

// CheckAddress - check that specified IP address can be pinged, and is available on specified port
func CheckAddress(address string) bool {
	p := fastping.NewPinger()                          // Creates new instance of fastping pinger
	ipAddress, err := net.ResolveIPAddr("ip", address) // Attempts to resolve IP
	p.AddIPAddr(ipAddress)                             // Adds resolved ip to pinger

	returnVal := false // Sets return value to false

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) { // On received packets from specified address
		returnVal = true // Sets return value
	}

	p.OnIdle = func() { // Timed out on address
		if returnVal != true { // Checks that address hasn't already been verified
			returnVal = false // Sets return value
		}
	}

	err = p.Run()
	if err != nil { // Checks for error
		returnVal = false // Sets return value
	}

	return returnVal // Return previously set return value
}

// GetCurrentTime - get current time in the UTC format
func GetCurrentTime() time.Time {
	return time.Now().UTC() // Returns current time in UTC
}
