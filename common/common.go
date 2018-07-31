package common

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
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
func CheckAddress(address string) error {
	p := fastping.NewPinger()
	ipAddress, err := net.ResolveIPAddr("ip", address)
	p.AddIPAddr(ipAddress)

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		fmt.Printf("IP %s tested successfully \n", addr.String())
	}
	p.OnIdle = func() {
		err = errors.New("Timed out with IP " + ipAddress.String() + "\n")
	}

	if err != nil {
		return err
	}

	err = p.Run()
	if err != nil {
		if strings.Contains(err.Error(), "operation not permitted") {
			return errors.New("operation requires root privileges")
		}

		return err
	}

	return nil
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
