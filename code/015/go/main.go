/* From Python to Go: Go: 015 - Basic SSH. */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
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

type Result struct {
	/* Struct to store command execution result. */
	Command   string
	Output    string
	Timestamp time.Time
}

type Device struct {
	/* Struct to interact with netowrk device. */
	Hostname    string `yaml:"hostname"`
	IpAddress   string `yaml:"ip_address"`
	Crendetials Crendetials
	Result      []Result
}

func (d *Device) executeCommand(c string) {
	/* Method to execute command */
	interactiveAuth := ssh.KeyboardInteractive(
		func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			answers := make([]string, len(questions))
			for i := range answers {
				answers[i] = (*d).Crendetials.Password
			}

			return answers, nil
		},
	)

	// Create a new SSH client
	sshClientConfig := &ssh.ClientConfig{
		User:            (*d).Crendetials.Username,
		Auth:            []ssh.AuthMethod{interactiveAuth},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%v:22", (*d).IpAddress), sshClientConfig)
	if err != nil {
		log.Fatalln("Failed to dial: ", err)
	}
	defer sshClient.Close()

	// Create session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalln("Failed to open the session: ", err)
	}
	defer session.Close()

	// Execute the command
	buffer := bytes.Buffer{}
	session.Stdout = &buffer
	if err := session.Run(c); err != nil {
		log.Fatalln("Failed to execute command: ", err)
	}

	// Update the result
	(*d).Result = append((*d).Result, Result{
		Command:   c,
		Output:    buffer.String(),
		Timestamp: time.Now(),
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

	// Execute commands
	for i := 0; i < len(*inventory); i++ {
		(*inventory)[i].Crendetials = sshCreds
		(*inventory)[i].executeCommand("show version")
	}

	// Print results
	for i := 0; i < len(*inventory); i++ {
		for j := 0; j < len((*inventory)[i].Result); j++ {
			fmt.Printf(
				"Command: %v\nOutput: %v\nTimestamp: %v\n",
				(*inventory)[i].Result[j].Command,
				(*inventory)[i].Result[j].Output,
				(*inventory)[i].Result[j].Timestamp,
			)
		}
	}
}
