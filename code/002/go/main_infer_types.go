package main

import (
	"fmt"
)

func main() {
	// Variables block
	username := "admin"
	ipAddress := "10.10.10.10"
	reconnectingAttempts := uint(3)
	isReconnecting := true

	// Printing block
	fmt.Printf(
		"You are connecting to %v with username %v, with reconnecting set to %v and %v attempts.\n",
		ipAddress,
		username,
		isReconnecting,
		reconnectingAttempts,
	)
}
