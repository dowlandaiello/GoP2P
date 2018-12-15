package common

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	upnp "github.com/NebulousLabs/go-upnp"
	fastping "github.com/tatsushid/go-fastping"
	"golang.org/x/crypto/sha3"
)

const (
	// NodeAvailableRep - global definition for reputation value of node availability
	NodeAvailableRep = 10

	// GoP2PTestnetID - GoP2P testing network identifier
	GoP2PTestnetID = 4519161392015

	// ConnectionDelimiter - GoP2P standard connection delimiter
	ConnectionDelimiter = byte('\f')

	// ProtobufPrefix - GoP2P standard protobuf message prefix
	ProtobufPrefix = "ProtoID"
)

var (
	// ExtIPProviders - preset macro defining list of available external IP checking services
	ExtIPProviders = []string{"http://checkip.amazonaws.com/", "http://icanhazip.com/", "http://www.trackip.net/ip", "http://bot.whatismyipaddress.com/", "https://ipecho.net/plain", "http://myexternalip.com/raw"}

	// GeneralTLSConfig - general global GoP2P TLS Config
	GeneralTLSConfig = &tls.Config{ // Init TLS config
		Certificates:       []tls.Certificate{getTLSCerts("GoP2PGeneral")},
		InsecureSkipVerify: true,
		ServerName:         "localhost"}

	// Silent - silence common.Println calls
	Silent = false
)

/*
	BEGIN EXPORTED METHODS:
*/

// Println - print
func Println(i interface{}) {
	if !Silent { // Check verbose mode
		fmt.Println(i) // Print
	}
}

// Print - print
func Print(i interface{}) {
	if !Silent { // Check verbose mode
		fmt.Print(i) // Print
	}
}

// Printf - print
func Printf(format string, i ...interface{}) {
	if !Silent {
		fmt.Printf(format, i...) // Print
	}
}

// DelaySeconds - wait until duration passed, return true once duration completed
func DelaySeconds(seconds uint) bool {
	startTime := time.Now() // Get current time

	for {
		if time.Now().Sub(startTime) >= time.Duration(seconds)*time.Second { // Check passed duration
			break // Break
		}
	}

	return true // Reached end
}

// GetCommonByteDifference - attempt to fetch most similar byte array in array of byte arrays
func GetCommonByteDifference(b [][]byte) ([]byte, error) {
	if len(b) == 0 { // Check for nil input
		return []byte{}, errors.New("nil input") // Return found error
	}

	if len(b) == 1 { // Check for single array
		return b[0], nil // Return single result
	}

	differences := []int{} // Init diff buffer

	lowest := 100    // Init lowest buffer
	lowestIndex := 0 // Init lowest index buffer

	sort.Slice(b, func(x, z int) bool { // Sort by length
		return len(b[x]) < len(b[z])
	})

	for x := 0; x != len(b); x++ { // Iterate through byte arrays
		diff := make([]int, len(b[x])) // Init diff buffer
		total := 0                     // Init total buffer

		for i := 0; i != len(b[x]); i++ { // Iterate through current byte array
			if x == len(b)-1 { // Check for end of slice
				diff[i] = int(b[x][i]) - int(b[x-1][i]) // Calculate difference
			} else {
				diff[i] = int(b[x][i]) - int(b[x+1][i]) // Calculate difference
			}

			total += diff[i] // Add value
		}

		if total/len(diff) < lowest { // Check value is lowest
			lowest = total / len(diff) // Set lowest
			lowestIndex = x            // Set index
		}

		differences = append(differences, total/len(diff)) // Append difference
	}

	return b[lowestIndex], nil // No error occurred, return nil
}

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
		Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt) // Print address meta
		Printf("IP %s tested successfully \n", addr.String())        // Print address meta
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

	return getNonNilInStringSlice(addresses) // Return valid address
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

	if len(strVal) < 30 { // Check for large string
		return strVal // Safe, return string
	}

	return strVal[0:30] // Unsafe, return sliced
}

// StringInSlice - check if value is in slice
func StringInSlice(s []string, v string) bool {
	for _, value := range s { // Iterate through values in slice
		if value == v { // Check for matching value
			return true // Value in slice
		}
	}

	return false // Value not in slice
}

// GenerateTLSCertificates - generate necessary TLS certificates, keys
func GenerateTLSCertificates(namePrefix string) error {
	_, certErr := os.Stat(fmt.Sprintf("%sCert.pem", namePrefix)) // Check for error reading file
	_, keyErr := os.Stat(fmt.Sprintf("%sKey.pem", namePrefix))   // Check for error reading file

	if os.IsNotExist(certErr) || os.IsNotExist(keyErr) { // Check for does not exist error
		privateKey, err := generateTLSKey(namePrefix) // Generate key

		if err != nil { // Check for errors
			return err // Return found error
		}

		err = generateTLSCert(privateKey, namePrefix) // Generate cert

		if err != nil { // Check for errors
			return err // Return found error
		}
	}

	return nil // No error occurred, return nil
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
			if len(*buffer) == 0 { // Double check IP not already determined
				*buffer = append(*buffer, "") // Set IP
				finished <- true              // Set finished
			}
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

// generateTLSKey - generates necessary TLS key
func generateTLSKey(namePrefix string) (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader) // Generate private key

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	marshaledPrivateKey, err := x509.MarshalECPrivateKey(privateKey) // Marshal private key

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: marshaledPrivateKey}) // Encode to memory

	err = ioutil.WriteFile(fmt.Sprintf("%sKey.pem", namePrefix), pemEncoded, 0644) // Write pem

	if err != nil { // Check for errors
		return nil, err // Return found error
	}

	return privateKey, nil // No error occurred, return nil
}

// generateTLSCert - generates necessary TLS cert
func generateTLSCert(privateKey *ecdsa.PrivateKey, namePrefix string) error {
	notBefore := time.Now() // Fetch current time

	notAfter := notBefore.Add(292 * (365 * (24 * time.Hour))) // Fetch 'deadline'

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)     // Init limit
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit) // Init serial number

	if err != nil { // Check for errors
		return err // Return found error
	}

	template := x509.Certificate{ // Init template
		SerialNumber: serialNumber, // Generate w/serial number
		Subject: pkix.Name{ // Generate w/subject
			Organization: []string{"localhost"}, // Generate w/org
		},
		NotBefore: notBefore, // Generate w/not before
		NotAfter:  notAfter,  // Generate w/not after

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // Generate w/key usage
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // Generate w/ext key
		BasicConstraintsValid: true,                                                         // Generate w/basic constraints
	}

	cert, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(privateKey), privateKey) // Generate certificate

	if err != nil { // Check for errors
		return err // Return found error
	}

	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert}) // Encode pem

	err = ioutil.WriteFile(fmt.Sprintf("%sCert.pem", namePrefix), pemEncoded, 0644) // Write cert file

	if err != nil { // Check for errors
		return err // Return found error
	}

	return nil // No error occurred, return nil
}

// getTLSCert - attempt to read TLS cert from current dir
func getTLSCerts(certPrefix string) tls.Certificate {
	GenerateTLSCertificates(certPrefix) // Generate certs

	cert, err := tls.LoadX509KeyPair(fmt.Sprintf("%sCert.pem", certPrefix), fmt.Sprintf("%sKey.pem", certPrefix)) // Load key pair

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	return cert // Return read certificates
}

func publicKey(privateKey interface{}) interface{} {
	switch k := privateKey.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func getNonNilInStringSlice(slice []string) (string, error) {
	for _, entry := range slice { // Iterate through entries
		if entry != "" { // Check for non-nil entry
			return entry, nil // Return valid entry
		}
	}

	return "", fmt.Errorf("couldn't find non-nil element in slice %v", slice) // Couldn't find valid address, return error
}

/*
	END INTERNAL METHODS
*/
