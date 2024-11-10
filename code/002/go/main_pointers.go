package main

import (
	"fmt"
)

func main() {
	// Declare variables
	username := new(string)
	ipAddress := new(string)
	reconnectingAttempts := new(uint)
	isReconnecting := new(bool)

	// Assign values
	*username = "admin"
	*ipAddress = "10.10.10.10"
	*reconnectingAttempts = 3
	*isReconnecting = true

	// Printing block
	fmt.Printf(
		"You are connecting to %v with username %v, with reconnecting set to %v and %v attempts.\n",
		ipAddress,
		username,
		isReconnecting,
		reconnectingAttempts,
	)

	fmt.Printf(
		"You are connecting to %v with username %v, with reconnecting set to %v and %v attempts.\n",
		*ipAddress,
		*username,
		*isReconnecting,
		*reconnectingAttempts,
	)
}
