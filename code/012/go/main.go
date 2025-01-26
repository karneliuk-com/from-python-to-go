/* From Python to Go: Go(Golang): 012 - User input. */

package main

// Import
import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/term"
)

// types
type CliFlags struct {
	Path string
}

type Credentials struct {
	Username string
	Password string
}

// Functions
func readArgs() CliFlags {
	/* Helper function to read CLI arguments. */

	// Prepare result
	result := CliFlags{}
	flag.StringVar(&result.Path, "p", "", "Path to the input file.")

	// Parse arguments
	flag.Parse()

	// Result
	return result
}

func loadFile(p string) string {
	/* Function to load a file. */
	bs, err := os.ReadFile(p)
	if err != nil {
		os.Exit(2)
	}

	// Result
	return string(bs)
}

func getCreds() Credentials {
	/* Helper function to get credentials */

	// Initialise result
	result := Credentials{}

	// Read Username
	fmt.Print("Username: ")
	_, err := fmt.Scanln(&result.Username)
	if err != nil {
		os.Exit(1)
	}

	// Read password
	fmt.Print("Password: ")
	bytepw, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		os.Exit(1)
	}
	result.Password = string(bytepw)

	// Return result
	return result
}

// Main
func main() {
	/* Main business logic */

	// Get arguments
	arg := readArgs()

	// load file
	if arg.Path != "" {
		fmt.Println(loadFile(arg.Path))

		// Exit if no path provided
	} else {
		os.Exit(3)
	}

	creds := getCreds()
	fmt.Printf("\n%+v\n", creds)
}
