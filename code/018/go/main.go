/* From Python to Go: Go: 018 - GNMI. */

package main

// Imports
import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmic/pkg/api"
	"google.golang.org/protobuf/encoding/prototext"
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
	Command string
	Config  struct {
		Path  string
		Value any
	}
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
	Port        uint   `yaml:"port"`
	Platform    string `yaml:"platform"`
	Crendetials Crendetials
	Result      []Result
}

func (d *Device) executeChange(i Instruction) {
	/* Method to execute command */

	// Create GNMI Target
	gnmiTarget, err := api.NewTarget(
		api.Name(d.Hostname),
		api.Address(fmt.Sprintf("%s:%d", d.IpAddress, d.Port)),
		api.Username(d.Crendetials.Username),
		api.Password(d.Crendetials.Password),
		api.SkipVerify(true),
	)
	if err != nil {
		log.Fatal("Cannot create GNMI Target: ", err)
	}

	// Create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create GNMI client
	err = gnmiTarget.CreateGNMIClient(ctx)
	if err != nil {
		log.Fatal("Cannot create GNMI Client: ", err)
	}
	defer gnmiTarget.Close()

	// Get state before change
	getReq, err := api.NewGetRequest(
		api.Path(i.Command),
		api.DataType("config"),
		api.Encoding("json_ietf"),
	)
	if err != nil {
		log.Fatal("Cannot create Get request: ", err)
	}
	beforeGetResponse, err := gnmiTarget.Get(ctx, getReq)
	if err != nil {
		log.Fatal("Cannot make a Get request: ", err)
	}
	beforeStruct := OpenConfigInterfaces{}
	err = json.Unmarshal(beforeGetResponse.Notification[0].Update[0].Val.GetJsonIetfVal(), &beforeStruct)
	if err != nil {
		log.Fatal("Cannot unmarshall JSON: ", err)
	}

	// Make change
	setReq, err := api.NewSetRequest(
		api.Update(
			api.Path(i.Config.Path),
			api.Value(i.Config.Value, "json_ietf"),
		),
	)
	if err != nil {
		log.Fatal("Cannot create Set request: ", err)
	}
	setResp, err := gnmiTarget.Set(ctx, setReq)
	if err != nil {
		log.Fatal("Cannot make a Set request: ", err)
	}
	log.Println(prototext.Format(setResp))

	// Get state after change
	afterGetResponse, err := gnmiTarget.Get(ctx, getReq)
	if err != nil {
		log.Fatal("Cannot make a Get request: ", err)
	}
	afterStruct := OpenConfigInterfaces{}
	err = json.Unmarshal(afterGetResponse.Notification[0].Update[0].Val.GetJsonIetfVal(), &afterStruct)
	if err != nil {
		log.Fatal("Cannot unmarshall JSON: ", err)
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
		Command: "/openconfig-interfaces:interfaces",
		Config: struct {
			Path  string
			Value any
		}{
			Path: "/openconfig-interfaces:interfaces",
			Value: map[string]any{
				"interface": []map[string]any{
					{
						"name": "Loopback 23",
						"config": map[string]any{
							"name":        "Loopback 23",
							"description": "Test-gnmi-golang-3",
						},
					},
				},
			},
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
