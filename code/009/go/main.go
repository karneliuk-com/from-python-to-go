/* From Python to Go: Go (Golang): 009 - Interfaces data type */

package main

// Imports
import (
	"fmt"
)

// Types
type Inventory []map[interface{}]interface{}

// Main
func main() {
	// Define inventory
	inventory := Inventory{
		{
			"name":      "leaf-01",
			"os":        "cisco-nxos",
			"ip":        "192.168.1.1",
			"port":      22,
			"latitude":  51.5120898,
			"longitude": -0.0030987,
			"active":    true,
		}, {
			"name":      "leaf-02",
			"os":        "arista-eos",
			"ip":        "192.168.1.2",
			"port":      830,
			"latitude":  51.5120427,
			"longitude": -0.0044585,
			"active":    true,
		}, {
			"name":      "spine-01",
			"ip":        "192.168.1.1",
			"port":      22,
			"latitude":  51.5112179,
			"longitude": -0.0048555,
			"active":    false,
		},
	}

	// Print the entire map and data type
	fmt.Printf("%+v, %T\n", inventory, inventory)

	// Print data types for each element
	for i := 0; i < len(inventory); i++ {
		fmt.Printf("%+v, %T\n", inventory[i], inventory[i])

		for k := range inventory[i] {
			fmt.Printf("%v=%v, %T\n", k, inventory[i][k], inventory[i][k])
		}
	}
}
