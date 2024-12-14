/* From Python to Go: Go (Golang): 007 - Classes and Structs" */

package main

// Imports
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Data types
type User struct {
	// Class to store user credentials
	username string
	password string
}

type Device struct {
	// Class to store device information
	hostname string
	port     uint64
	ip       string
	nos      string
}

// Class to store inventory information
type Inventory []Device

// Aux functions
func getCredentials() User {
	// Function to retrieve credentials from the environment
	blocks := strings.Split(os.Getenv("AUTOMATION_CREDS"), ",")
	return User{
		blocks[0],
		blocks[1],
	}
}

func getInventory() *Inventory {
	// Create an empty list to store devices
	result := &Inventory{}

	// Loop through the environment variables
	for _, kv := range os.Environ() {
		// Check if the key starts with AUTOMATION_DEVICE_
		if strings.Contains(kv, "AUTOMATION_DEVICE_") {
			// Split the value by comma and create a new device object
			blocks := strings.Split(strings.Split(kv, "=")[1], ",")

			devicePort, err := strconv.ParseUint(blocks[1], 10, 64)
			if err != nil {
				fmt.Printf("Got error when converting string to uint: %v\n", err)
				os.Exit(1)
			}

			*result = append(*result, Device{
				blocks[0],
				devicePort,
				blocks[3],
				blocks[2],
			})
		}
	}

	// Result
	return result
}

// Main function
func main() {
	/* Main business logic */

	// Get the credentials
	user := getCredentials()

	// Print credentails
	fmt.Printf("%+v\n", user)

	// Get inventory
	inventory := getInventory()

	// Print inventory memory address
	fmt.Printf("Memory address of inventory: %v\n", &inventory)

	// Print inventory content
	fmt.Printf("%+v\n", *inventory)
}
