/* From Python to Go: Go: 017 - NETCONF. */

package main

// Imports
import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/options"
	"gopkg.in/yaml.v3"
)

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
	Command OpenConfigInterfaces
	Config  OpenConfigInterfaces
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

	// Create XML for filter
	// xmlCommandFilter, err := xml.Marshal(i.Command)
	// if err != nil {
	// 	log.Fatalln("Cannot create XML message ", err)
	// }

	// Get netowrk driver
	dr, err := netconf.NewDriver(
		(*d).IpAddress,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername(d.Crendetials.Username),
		options.WithAuthPassword(d.Crendetials.Password),
	)
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
	before, err := dr.GetConfig("running")
	if err != nil {
		log.Fatalln("failed to send command; error: ", err)
	}
	beforeStruct := RPCResponse{}
	err = xml.Unmarshal((*before).RawResult, &beforeStruct)
	if err != nil {
		log.Panic("Cannot parse received response: ", err)
	}

	// Apply change
	configXmlBs, err := xml.Marshal(i.Config)
	if err != nil {
		log.Fatalln("Cannot convert config to XML: ", err)
	}
	configXmlStr := "<config>" + string(configXmlBs) + "</config>"

	changeResponse, err := dr.EditConfig("candidate", configXmlStr)
	if err != nil {
		log.Fatalln("failed to send config; error: ", err)
	} else if changeResponse.Failed != nil {
		log.Fatalln("Return error from device during config; error: ", err)
	}

	commitResponse, err := dr.Commit()
	if err != nil {
		log.Fatalln("failed to commit config; error: ", err)
	} else if commitResponse.Failed != nil {
		log.Fatalln("return error from device during commit; error: ", err)
	}

	// Get state after change
	after, err := dr.GetConfig("running")
	if err != nil {
		log.Fatalln("failed to send command; error: ", err)
	}
	afterStruct := RPCResponse{}
	err = xml.Unmarshal((*after).RawResult, &afterStruct)
	if err != nil {
		log.Panic("Cannot parse received response: ", err)
	}

	// Diff
	diff := cmp.Diff(beforeStruct, afterStruct)

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
		Command: OpenConfigInterfaces{
			XMLName: xml.Name{
				Space: "http://openconfig.net/yang/interfaces",
				Local: "interfaces",
			},
		},
		Config: OpenConfigInterfaces{
			XMLName: xml.Name{
				Space: "http://openconfig.net/yang/interfaces",
				Local: "interfaces",
			},
			Interface: make([]OpenConfigInterface, 0),
		},
	}
	instruction.Config.Interface = append(instruction.Config.Interface, OpenConfigInterface{
		Name: "Loopback 23",
		Config: struct {
			Name        string "xml:\"name,omitempty\""
			Description string "xml:\"description,omitempty\""
		}{
			Name:        "Loopback 23",
			Description: "Test-netconf-golang-2",
		},
	})

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
