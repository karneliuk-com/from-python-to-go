/* From Python to Go: Go (Golang): 011 - Parsing XML/JSON/YAML files */
package main

// Import
import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Types
type InventoryItem struct {
	/* Device -- Inventory Item */
	Name      string  `xml:"name" json:"name" yaml:"name"`
	OS        string  `xml:"os" json:"os" yaml:"os"`
	IP        string  `xml:"ip" json:"ip" yaml:"ip"`
	Port      int64   `xml:"port" json:"port" yaml:"port"`
	Latitude  float64 `xml:"latitude" json:"latitude" yaml:"latitude"`
	Longitude float64 `xml:"longitude" json:"longitude" yaml:"longitude"`
	Active    bool    `xml:"active" json:"active" yaml:"active"`
}
type Inventory struct {
	/* Inventory of all devices */
	Devices []InventoryItem `xml:"devices" json:"devices" yaml:"devices"`
}

// Aux functions
func loadInventory(p string) Inventory {
	/* Function to load inventory */

	// Load file
	bs, err := os.ReadFile(p)
	if err != nil {
		fmt.Printf("Cannot open file %v: %v\n", p, err)
		os.Exit(1)
	}

	// Define result
	result := Inventory{}

	// Find importer
	reXML := regexp.MustCompile(`^.+\.xml$`)
	reJSON := regexp.MustCompile(`^.+\.json$`)
	reYAML := regexp.MustCompile(`^.+\.ya?ml$`)

	// XML parsing
	if reXML.MatchString(p) {
		err := xml.Unmarshal(bs, &result)
		if err != nil {
			fmt.Printf("Cannot parse XML data: %v\n", err)
		}
		// JSON parsing
	} else if reJSON.MatchString(p) {
		err := json.Unmarshal(bs, &result)
		if err != nil {
			fmt.Printf("Cannot parse JSON data: %v\n", err)
		}
		// YAML parsing
	} else if reYAML.MatchString(p) {
		err := yaml.Unmarshal(bs, &result)
		if err != nil {
			fmt.Printf("Cannot parse YAML data: %v\n", err)
		}
	} else {
		fmt.Printf("Unknown file format: %v\n", p)
	}

	// Return result
	return result
}

// Main
func main() {
	/* Main busines logic */

	// Check that file is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./app <file>")
		os.Exit(1)
	}

	// Load inventory
	inv := loadInventory(os.Args[1])

	// Print inventory
	fmt.Printf("%+v\n", inv)
}
