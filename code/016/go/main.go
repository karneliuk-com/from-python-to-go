/* From Python to Go: Go: 016 - Advanced SSH. */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
	"gopkg.in/yaml.v3"
)

// Imports

// Types and Receivers
type Arguments struct {
	/* Class to starte CLI arguments */
	Inventory string
}

type Crendetials struct {
	/* Struct to store credentials. */
	Username string
	Password string
}

type Instruction struct {
	Command string
	Config  []string
}

type Result struct {
	/* Struct to store command execution result. */
	Instruction Instruction
	Diff        string
	Timestamp   time.Time
}

type Device struct {
	/* Struct to interact with netowrk device. */
	Hostname    string `yaml:"hostname"`
	IpAddress   string `yaml:"ip_address"`
	Platform    string `yaml:"platform"`
	Crendetials Crendetials
	Result      []Result
}

func (d *Device) executeChange(i Instruction) {
	/* Method to execute command */

	// Get platform
	p, err := platform.NewPlatform(
		(*d).Platform,
		(*d).IpAddress,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername(d.Crendetials.Username),
		options.WithAuthPassword(d.Crendetials.Password),
	)
	if err != nil {
		log.Fatalln("failed to create platform; error: ", err)
	}

	// Get netowrk driver
	dr, err := p.GetNetworkDriver()
	if err != nil {
		log.Fatalln("failed to fetch network driver from the platform; error ", err)
	}

	// Open session
	err = dr.Open()
	if err != nil {
		log.Fatalln("failed to open driver; error: ", err)
	}
	defer dr.Close()

	// Get change before start
	before, err := dr.SendCommand(i.Command)
	if err != nil {
		log.Fatalln("failed to send command; error: ", err)
	}

	// Apply change
	_, err = dr.SendConfigs(i.Config)
	if err != nil {
		log.Fatalln("failed to send config; error: ", err)
	}

	// Get state after change
	after, err := dr.SendCommand(i.Command)
	if err != nil {
		log.Fatalln("failed to send command; error: ", err)
	}

	// Diff
	diff := cmp.Diff(strings.Split(before.Result, "\n"), strings.Split(after.Result, "\n"))

	// Update the result
	(*d).Result = append((*d).Result, Result{
		Instruction: i,
		Diff:        diff,
		Timestamp:   time.Now(),
	})
}

// Functions
func readArgs() Arguments {
	/* Helper function to read CLI arguments */
	result := Arguments{}

	flag.StringVar(&result.Inventory, "i", "", "Path to the inventory file")

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

func getCredentials() Crendetials {
	/* Function to get credentials. */
	return Crendetials{
		Username: os.Getenv("AUTOMATION_USER"),
		Password: os.Getenv("AUTOMATION_PASS"),
	}
}

// Main
func main() {
	/* Core logic */
	// Read CLI arguments
	cliArgs := readArgs()

	// Get credentials
	sshCreds := getCredentials()

	// Load inventory
	inventory := loadInventory(cliArgs.Inventory)

	// Config
	instruction := Instruction{
		Command: "show interfaces description",
		Config: []string{
			"interface Loopback 23",
			"description Go_Test_2",
		},
	}

	// Execute commands
	for i := 0; i < len(*inventory); i++ {
		(*inventory)[i].Crendetials = sshCreds
		(*inventory)[i].executeChange(instruction)
	}

	// Print results
	for i := 0; i < len(*inventory); i++ {
		for j := 0; j < len((*inventory)[i].Result); j++ {
			fmt.Printf(
				"Config: %v\nImpact: %v\nTimestamp: %v\n",
				(*inventory)[i].Result[j].Instruction.Config,
				(*inventory)[i].Result[j].Diff,
				(*inventory)[i].Result[j].Timestamp,
			)
		}
	}
}
