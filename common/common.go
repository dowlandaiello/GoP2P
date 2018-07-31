package common

import (
	"net"
	"net/http"
	"time"

	upnp "github.com/NebulousLabs/go-upnp"
	"github.com/mitsukomegumi/GoP2P/fastping"
)

const (
	// NodeAvailableRep - global definition for value of node availability
	NodeAvailableRep = 10
)

/*
	BEGIN EXPORTED METHODS:
*/

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

// GetExtIPAddrWithUpNP - retrieve the external IP address of the current machine via upnp
func GetExtIPAddrWithUpNP() (string, error) {
	// connect to router
	d, err := upnp.Discover()
	if err != nil { // Check for errors
		return "", err // return error
	}

	// discover external IP
	ip, err := d.ExternalIP()
	if err != nil { // Check for errors
		return "", err // return error
	}
	return ip, nil
}

// GetExtIPAddrWithoutUpNP - retrieve the external IP address of the current machine w/o upnp
func GetExtIPAddrWithoutUpNP() (string, error) {
	ip := make([]byte, 100) // Create IP buffer

	resp, err := http.Get("http://checkip.amazonaws.com/") // Attempt to check IP via aws
	if err != nil {                                        // Check for errors
		return "", err // Return error
	}

	defer resp.Body.Close()     // Close connection
	_, err = resp.Body.Read(ip) // Read IP

	if err != nil { // Check for errors
		return "", err // Return error
	}

	return string(ip[:len(ip)]), nil // Return ip
}

// GetCurrentTime - get current time in the UTC format
func GetCurrentTime() time.Time {
	return time.Now().UTC() // Returns current time in UTC
}

/*
	END EXPORTED METHODS
*/
