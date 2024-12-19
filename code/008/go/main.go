/* From Python to Go: Go (Golang): 008 - Object-oriented programming */

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

func (u *User) getCredentials() {
	// Function to retrieve credentials from the environment
	blocks := strings.Split(os.Getenv("AUTOMATION_CREDS"), ",")
	(*u).username = blocks[0]
	(*u).password = blocks[1]
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

func (i *Inventory) populate() {
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

			*i = append(*i, Device{
				blocks[0],
				devicePort,
				blocks[3],
				blocks[2],
			})
		}
	}
}

// Main function
func main() {
	/* Main business logic */

	// Get the credentials
	user := User{}
	user.getCredentials()

	// Print credentails
	fmt.Printf("%+v\n", user)

	// Get inventory
	inventory := Inventory{}
	inventory.populate()

	// Print inventory memory address
	fmt.Printf("Memory address of inventory: %v\n", &inventory)

	// Print inventory content
	fmt.Printf("%+v\n", inventory)
}
