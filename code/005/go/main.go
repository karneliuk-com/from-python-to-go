/* From Python to Go: Python: 005 - Code Flow Control */

package main

import "fmt"

// Aux functions
func getData(ipAddress string) []string {
	/* Function that pretends to connect to network device and collect some output */

	// Ensure there is IP address to connect
	if ipAddress == "" {
		return []string{"NOT OK", "There is no IP address provided"}
	}

	// Return some mock data
	return []string{"OK", "Some raw data"}
}

func parseData(parser, data string) []string {
	/* Function that pretends to parse the output from network device depeinding on its operating system type */

	switch parser {
	case "cisco-nxos":
		return []string{"OK", fmt.Sprintf("parsed: %v\n", data)}
	case "arista-eos":
		return []string{"OK", fmt.Sprintf("parsed: %v\n", data)}
	default:
		return []string{"NOT-OK", "There is no parser available"}
	}
}

// Main
func main() {
	// Define inventory
	inventory := [][]string{
		{"leaf-01", "cisco-nxos", "192.168.1.1"},
		{"leaf-02", "arista-eos", "192.168.1.2"},
		{"spine-01", "cisco-nxos", ""},
		{"spine-02", "arosta-eos", "192.168.1.12"},
	}

	// Loop through all network devices
	for _, device := range inventory {
		// Collect data for each network device
		collectedData := getData(device[2])

		var parsedData []string

		// Do parsing if data is collected
		if collectedData[0] == "OK" {
			parsedData = parseData(device[1], collectedData[1])
		} else {
			fmt.Printf("Collecting data from %v is not successful\n", device[0])
			continue
		}

		// Print results
		if parsedData[0] == "OK" {
			fmt.Printf("Successfully collected and parsed data for %v\n", device[0])
		} else {
			fmt.Printf("Successfully collected but NOT parsed data for %v\n", device[0])
		}
	}
}
