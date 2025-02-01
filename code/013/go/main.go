/* From Python to Go: Go(Golang): 013 -  Exception handling. */

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
		fmt.Println(err)
		os.Exit(1)
	}

	// Result
	return string(bs)
}

// Here is an important trick, for output provide (var_name data_type)
func getCreds() (result Credentials) {
	/* Helper function to get credentials */

	// Catching error
	defer func() {
		/* Helper function for recovery */
		r := recover()

		if r != nil {
			fmt.Printf("Recovering from '%v'\n", r)
		}
	}()

	// Read Username
	fmt.Print("Username: ")
	_, err := fmt.Scanln(&result.Username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read password
	fmt.Print("Password: ")
	bytepw, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result.Password = string(bytepw)

	// If password isn't provided, throw an exception
	if len(result.Password) == 0 {
		panic("No password is provided!")
	}

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
