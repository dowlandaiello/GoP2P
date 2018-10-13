package common

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	upnp "github.com/NebulousLabs/go-upnp"
	fastping "github.com/tatsushid/go-fastping"
)

const (
	// NodeAvailableRep - global definition for reputation value of node availability
	NodeAvailableRep = 10
)

var (
	// ExtIPProviders - preset macro defining list of available external IP checking services
	ExtIPProviders = []string{"http://checkip.amazonaws.com/"}
)

/*
	BEGIN EXPORTED METHODS:
*/

// SeedAddress - generated shard address from seed
func SeedAddress(seed string, shardID string) (string, error) {
	if len(seed) == 0 || len(shardID) == 0 {
		return "", errors.New("invalid input") // Return found error
	}
	for i := range seed { // Iterate over seed
		if i%3 == 0 { // Check is third
			seed = seed[:i] + string(shardID[i]) + seed[i+1:] // Replace with ID at string
		}
	}

	return seed, nil // Return seed
}

// ParseStringMethodCall - attempt to parse string as method call, returning receiver, method name, and params
func ParseStringMethodCall(input string) (string, string, []string, error) {
	if input == "" { // Check for errors
		return "", "", []string{}, errors.New("nil input") // Return found error
	} else if !strings.Contains(input, "(") || !strings.Contains(input, ")") {
		input = input + "()" // Fetch receiver methods
	}

	if !strings.Contains(input, ".") { // Check for nil receiver
		return "", "", []string{}, errors.New("invalid method " + input) // Return found error
	}

	method := strings.Split(strings.Split(input, "(")[0], ".")[1] // Fetch method

	receiver := StringFetchCallReceiver(input) // Fetch receiver

	params := []string{} // Init buffer

	if !strings.Contains(input, "()") { // Check for nil params
		params, _ = ParseStringParams(input) // Fetch params
	}

	return receiver, method, params, nil // No error occurred, return parsed method+params
}

// ParseStringParams - attempt to fetch string parameters from (..., ..., ...) style call
func ParseStringParams(input string) ([]string, error) {
	if input == "" { // Check for errors
		return []string{}, errors.New("nil input") // Return found error
	}

	parenthesesStripped := StringStripParentheses(StringStripReceiverCall(input)) // Strip parentheses

	params := strings.Split(parenthesesStripped, ", ") // Split by ', '

	return params, nil // No error occurred, return split params
}

// ConvertStringToReflectValues - convert string to []reflect.Value
func ConvertStringToReflectValues(inputs []string) []reflect.Value {
	values := []reflect.Value{} // Init buffer

	for input := range inputs {
		values = append(values, reflect.ValueOf(input)) // Add reflect value
	}

	return values
}

// StringStripReceiverCall - strip receiver from string method call
func StringStripReceiverCall(input string) string {
	return "(" + strings.Split(input, "(")[1] // Split
}

// StringStripParentheses - strip parantheses from string
func StringStripParentheses(input string) string {
	leftStripped := strings.Replace(input, "(", "", -1) // Strip left parent

	return strings.Replace(leftStripped, ")", "", -1) // Return right stripped
}

// StringFetchCallReceiver - attempt to fetch receiver from string, as if it were a x.y(..., ..., ...) style method call
func StringFetchCallReceiver(input string) string {
	return strings.Split(strings.Split(input, "(")[0], ".")[0] // Return split string
}

// Forever - prevent thread from closing
func Forever() {
	for {
		time.Sleep(time.Second)
	}
}

// CheckAddress - check that specified IP address can be pinged, and is available on specified port
func CheckAddress(address string) error {
	if address == "" { // Check for nil address
		return errors.New("nil address") // Return error
	}

	p := fastping.NewPinger() // Create new pinger

	ipAddress, err := net.ResolveIPAddr("ip", address) // Resolve address

	if err != nil {
		return err
	}

	p.AddIPAddr(ipAddress) // Add address to pinger

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) { // On correct address
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt) // Print address meta
		fmt.Printf("IP %s tested successfully \n", addr.String())        // Print address meta
	}
	p.OnIdle = func() { // On address timeout
		err = errors.New("Timed out with IP " + ipAddress.String() + "\n") // Assign meta to error
	}

	if err != nil { // Check for error
		return err // Return found error
	}

	err = p.Run()   // Assign to error
	if err != nil { // Check for errors
		return err // Return error
	}

	return nil // No error found, return nil
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
	resp, err := http.Get("http://checkip.amazonaws.com/") // Attempt to check IP via aws
	if err != nil {                                        // Check for errors
		return "", err // Return error
	}

	defer resp.Body.Close() // Close connection

	ip, err := ioutil.ReadAll(resp.Body) // Read address

	if err != nil { // Check for errors
		return "", err // Return error
	}

	stringVal := string(ip[:]) // Fetch string value

	return strings.TrimSpace(stringVal), nil // Return ip
}

// GetCurrentTime - get current time in the UTC format
func GetCurrentTime() time.Time {
	return time.Now().UTC() // Returns current time in UTC
}

// GetCurrentDir - returns current execution directory
func GetCurrentDir() (string, error) {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil { // Check for errors
		return "", err
	}
	return currentDir, nil
}

// SHA256 - hash specified byte array
func SHA256(b []byte) string {
	hash := sha256.Sum256(b)                          // Hash it
	return base64.StdEncoding.EncodeToString(hash[:]) // Return it
}

/*
	END EXPORTED METHODS
*/
