/* From Python to Go: Go (Golang): 010 - Text files */

package main

// Import
import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Aux functions
func loadFile(f string) string {
	/* Helper function to read file */

	// Read file
	bs, err := os.ReadFile(f)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Return data
	return string(bs)
}
func saveFile(f string, c []string) bool {
	/* Helper function to write to file */

	// Prepare data for write
	data := []byte(strings.Join(c, "\n"))

	// Write to file
	err := os.WriteFile(f, data, 0644)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return false
	}

	// Return result
	return true
}

// Main
func main() {
	// Get paths
	fileToOpen := "../data/file.txt"
	fileToSave := "../data/output.txt"

	// Read file
	data := loadFile(fileToOpen)

	// Print the raw text file
	fmt.Println(data)

	// File is a multiline string, so split it to lines
	newData := []string{}
	for ind, line := range strings.Split(data, "\n") {
		fmt.Printf("line %03d: %v\n", ind, line)

		// Make FQDN
		re := regexp.MustCompile(`^hostname:\s+.*$`)
		if re.MatchString(line) && !strings.Contains(line, "network.karneliuk.com") {
			line += ".network.karneliuk.com"
		}

		// Copy line to new output
		newData = append(newData, line)
	}

	// Save result to file
	r := saveFile(fileToSave, newData)
	if r {
		fmt.Println("File is saved successfully.")
	} else {
		fmt.Println("File is NOT saved successfully.")
	}
}
