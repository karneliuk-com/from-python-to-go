/* From Python to Go: Go: 014 - Templating configuration. */

package main

// Imports
import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

// Types
type IPAddress struct {
	/* Class to store IP address data. */
	Address string `yaml:"address"`
	Prefix  int    `yaml:"prefix"`
}
type Interface struct {
	/* Class to store interface data. */
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	IP4         IPAddress `yaml:"ip4"`
	Enabled     bool      `yaml:"enabled"`
}
type Device struct {
	/* Class to store credentials. */
	Hostname   string      `yaml:"hostname"`
	Interfaces []Interface `yaml:"interfaces"`
}
type Arguments struct {
	/* Class to starte CLI arguments */
	Data     string
	Template string
}

// Functions
func readArgs() Arguments {
	/* Helper function to read CLI arguments */
	result := Arguments{}

	flag.StringVar(&result.Data, "d", "", "Path to the input file")
	flag.StringVar(&result.Template, "t", "", "Path to the template.")

	flag.Parse()

	return result
}

func loadInventory(p string) *[]Device {
	/* Function to load inventory data. */

	// Open file
	bs, err := os.ReadFile(p)
	if err != nil {
		fmt.Println("Get error ", err)
		os.Exit(1)
	}

	// Load inventory
	result := &[]Device{}

	err = yaml.Unmarshal(bs, result)
	if err != nil {
		fmt.Println("Get error ", err)
		os.Exit(1)
	}

	// Return result
	return result
}

func loadTemplate(p string) *template.Template {
	/* Helper function to load template. */

	// Load template
	templ, err := template.New(p).ParseFiles(p)
	if err != nil {
		fmt.Println("Get error ", err)
		os.Exit(1)
	}

	// Return result
	return templ
}

func createConfiguration(d *[]Device, t *template.Template) bool {
	/* Function to create configuration files. */

	// Create output directory
	err := os.MkdirAll("output", 0777)
	if err != nil {
		fmt.Println("Get error ", err)
		return false
	}

	// Render template
	for i := 0; i < len(*d); i++ {
		// Create file
		f, err := os.Create("output/" + (*d)[i].Hostname + ".txt")
		if err != nil {
			fmt.Println("Get error ", err)
			f.Close()
			return false
		}

		// Render template
		t.Execute(f, (*d)[i])

		// Close file
		f.Close()
	}

	return true
}

// Main
func main() {
	// Get arguments
	args := readArgs()

	// Load inventory
	inventory := loadInventory(args.Data)

	// Load template
	templ := loadTemplate(args.Template)

	// Create configuration
	if createConfiguration(inventory, templ) {
		fmt.Println("Configuration files created.")
	} else {
		fmt.Println("Something went wrong.")
	}
}
