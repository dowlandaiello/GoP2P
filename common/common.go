package common

import (
	"encoding/hex"
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
	"golang.org/x/crypto/sha3"
)

const (
	// NodeAvailableRep - global definition for reputation value of node availability
	NodeAvailableRep = 10

	// GoP2PTestNetID - GoP2P testing network identifier
	GoP2PTestNetID = 4519161392015

	// ConnectionDelimiter - GoP2P standard connection delimiter
	ConnectionDelimiter = byte('\f')
)

var (
	// ExtIPProviders - preset macro defining list of available external IP checking services
	ExtIPProviders = []string{"http://checkip.amazonaws.com/", "http://icanhazip.com/", "http://www.trackip.net/ip", "http://bot.whatismyipaddress.com/", "https://ipecho.net/plain", "http://myexternalip.com/raw"}
)

/*
	BEGIN EXPORTED METHODS:
*/

// SeedAddress - generated shard address from seeds
func SeedAddress(seeds []string, shardID string) (string, error) {
	if len(seeds) == 0 || len(shardID) == 0 || len(shardID) < len(seeds) { // Check for invalid input
		return "", errors.New("invalid input") // Return found error
	}

	seed := shardID[0:5] + "::" + strings.Join(seeds, ":") // Set seed

	return seed, nil // Return seed
}

// ParseShardAddress - attempt to fetch node addresses from shard address
func ParseShardAddress(address string) ([]string, error) {
	if address == "" || !strings.Contains(address[5:7], "::") { // Check for nil input
		return []string{}, errors.New("invalid input") // Return found error
	}

	addresses := strings.Split(address[7:(len(address)-1)], ":") // Split

	return addresses, nil // Return addresses
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

// GetExtIPAddrWithUPnP - retrieve the external IP address of the current machine via upnp
func GetExtIPAddrWithUPnP() (string, error) {
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

// GetExtIPAddrWithoutUPnP - retrieve the external IP address of the current machine w/o upnp
func GetExtIPAddrWithoutUPnP() (string, error) {
	addresses := []string{} // Init address buffer

	finished := make(chan bool) // Init finished

	for _, provider := range ExtIPProviders { // Iterate through providers
		go getIPFromProviderAsync(provider, &addresses, finished) // Fetch IP
	}

	<-finished // Wait until finished

	close(finished) // Close channel

	return getNonNilInStringSlice(addresses), nil // Return valid address
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

// Sha3 - hash specified byte array
func Sha3(b []byte) string {
	hash := sha3.New256() // Hash

	hash.Write(b) // Write hash

	return hex.EncodeToString(hash.Sum(nil)) // Return hash
}

// SafeSlice - fetch sliced byte array input (max 20 chars)
func SafeSlice(b []byte) string {
	strVal := string(b) // Convert to string

	if len(strVal) < 20 { // Check for large string
		return strVal // Safe, return string
	}

	return strVal[0:20] // Unsafe, return sliced
}

/*
	END EXPORTED METHODS
*/

/*
	BEGIN INTERNAL METHODS:
*/

func getIPFromProvider(provider string) (string, error) {
	resp, err := http.Get(provider) // Attempt to check IP via provider

	if err != nil { // Check for errors
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

func getIPFromProviderAsync(provider string, buffer *[]string, finished chan bool) {
	if len(*buffer) == 0 { // Check IP not already determined
		resp, err := http.Get(provider) // Attempt to check IP via provider

		if err != nil { // Check for errors
			fmt.Println(err.Error()) // Log error
		} else {
			defer resp.Body.Close() // Close connection

			ip, _ := ioutil.ReadAll(resp.Body) // Read address

			stringVal := string(ip[:]) // Fetch string value

			if len(*buffer) == 0 { // Double check IP not already determined
				*buffer = append(*buffer, strings.TrimSpace(stringVal)) // Set ip

				finished <- true // Set finished
			}
		}
	}
}

func getNonNilInStringSlice(slice []string) string {
	for _, entry := range slice { // Iterate through entries
		if entry != "" { // Check for non-nil entry
			return entry // Return valid entry
		}
	}

	return "" // Couldn't find valid address, return nil
}

/*
	END INTERNAL METHODS
*/
