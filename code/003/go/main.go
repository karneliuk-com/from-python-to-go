/*From Python to Go: Go: 003 - imports and function*/
package main

// Import part
import (
	"fmt"
	"os"
)

// Aux functions
func getUsername() string {
	/* Helper function to get username */
	return os.Getenv("AUTOMATION_USERNAME")
}

func getPassword() string {
	/* Helper function to get password */
	return os.Getenv("AUTOMATION_PASSWORD")
}

// Main
func main() {
	// get variables
	username := getUsername()
	password := getPassword()

	fmt.Printf("username=%v, password=%v\n", username, password)
}
