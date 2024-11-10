package main

import (
	"fmt"
)

func main() {
	// Declare variables
	var username string
	var ipAddress string
	var reconnectingAttempts uint
	var isReconnecting bool

	// Assign values
	username = "admin"
	ipAddress = "10.10.10.10"
	reconnectingAttempts = 3
	isReconnecting = true

	// Printing block
	fmt.Printf(
		"You are connecting to %v with username %v, with reconnecting set to %v and %v attempts.\n",
		ipAddress,
		username,
		isReconnecting,
		reconnectingAttempts,
	)
}
