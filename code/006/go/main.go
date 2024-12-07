/* From Python to Go: Python: 006 - Dictionaries and Maps */
package main

import (
	"fmt"
)

func main() {
	// Define inventory
	inventory := []map[string]string{
		{
			"name": "leaf-01",
			"os":   "cisco-nxos",
			"ip":   "192.168.1.1",
		}, {
			"name": "leaf-02",
			"os":   "arista-eos",
			"ip":   "192.168.1.2",
		}, {
			"name": "spine-01",
			"ip":   "192.168.1.1",
		},
	}

	// Loop through all network devices
	for _, d := range inventory {
		// Print the device data
		fmt.Println(d)

		// Print the hostname
		fmt.Printf("Hostname: %v\n", d["name"])

		// Add the OS key if it is missing
		if _, ok := d["os"]; !ok {
			d["os"] = ""
		}
		fmt.Println(d)

		// Add new key-value pair
		d["location"] = "DC1"
		fmt.Println(d)

		// Remove the IP key
		delete(d, "ip")
		fmt.Println(d)

		// Go through all keys and values
		for k, v := range d {
			fmt.Printf("%v: %v\n", k, v)
		}
	}

	fmt.Println(inventory)
}
