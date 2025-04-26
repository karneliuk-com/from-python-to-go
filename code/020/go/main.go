/* From Python to Go: Go: 020 - Concurency. */

package main

// Imports
import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmic/pkg/api"
	"google.golang.org/protobuf/encoding/prototext"
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

type InventoryCredentials struct {
	/* Struct to store inventory credentails */
	Url   string
	Token string
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
func loadInventory(iC InventoryCredentials) *[]Device {
	/* Function to load inventory data. */

	// Create HTTP client
	hclient := &http.Client{}

	// Prepare request
	NetboxRequest, err := http.NewRequest("GET", iC.Url+"/api/dcim/devices/", nil)
	if err != nil {
		fmt.Println("Error during preparing HTTP Request ", err)
		os.Exit(1)
	}

	// Set headers
	NetboxRequest.Header.Add("Authorization", fmt.Sprintf("Token %s", iC.Token))

	// Set URL params
	q := NetboxRequest.URL.Query()
	q.Add("site", "ka-blog")
	NetboxRequest.URL.RawQuery = q.Encode()

	// Get data
	resp, err := hclient.Do(NetboxRequest)
	if err != nil {
		fmt.Println("Erorr during executing HTTP query ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error during reading body of HTTP response ", err)
		os.Exit(1)
	}

	td := NetboxDcimDevices{}
	err = json.Unmarshal(body, &td)
	if err != nil {
		fmt.Println("Error during parsing JSON ", err)
		os.Exit(1)
	}

	// Load inventory
	result := &[]Device{}

	for _, v := range td.Results {
		*result = append(*result, Device{
			Hostname:  v.Name,
			Platform:  v.Platform.Slug,
			IpAddress: strings.Split(v.PrimaryIp4.Address, "/")[0],
			Port:      uint(v.CustomFields.GnmiPort),
		})
	}

	// Return result
	return result
}

func getCredentials() (Crendetials, InventoryCredentials) {
	/* Function to get credentials. */
	return Crendetials{
			Username: os.Getenv("AUTOMATION_USER"),
			Password: os.Getenv("AUTOMATION_PASS"),
		},
		InventoryCredentials{
			Url:   os.Getenv("AUTOMATION_INVENTORY_URL"),
			Token: os.Getenv("AUTOMATION_INVENTORY_TOKEN"),
		}
}

// Main
func main() {
	/* Core logic */
	// Get credentials
	sshCreds, invCreds := getCredentials()

	// Load inventory
	inventory := loadInventory(invCreds)

	// Config
	instruction := Instruction{
		Command: "/interfaces",
		Config: struct {
			Path  string
			Value any
		}{
			Path: "/interfaces",
			Value: map[string]any{
				"interface": []map[string]any{
					{
						"name": "Loopback 23",
						"config": map[string]any{
							"name":        "Loopback 23",
							"description": "Test-gnmi-golang-23",
						},
					},
				},
			},
		},
	}

	// Create communication channel
	c := make(chan Device)

	// Execute commands
	for i := 0; i < len(*inventory); i++ {
		// Set credentals
		(*inventory)[i].Crendetials = sshCreds

		// Launch goroutines
		go func(d Device, ins Instruction, c chan<- Device) {
			// Execute task
			d.executeChange(ins)

			// Send device back
			c <- d
		}((*inventory)[i], instruction, c)
	}

	// Collect results
	iventory_with_results := make([]Device, 0)
	for i := 0; i < len(*inventory); i++ {
		iventory_with_results = append(iventory_with_results, <-c)
	}

	// Print results
	for i := 0; i < len(iventory_with_results); i++ {
		for j := 0; j < len((iventory_with_results)[i].Result); j++ {
			fmt.Printf(
				"Config: %v\nImpact: %v\nTimestamp: %v\n",
				(iventory_with_results)[i].Result[j].Instruction.Config,
				(iventory_with_results)[i].Result[j].Diff,
				(iventory_with_results)[i].Result[j].Timestamp,
			)
		}
	}
}
