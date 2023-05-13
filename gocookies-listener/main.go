package main

import "log"

func main() {
	listener, err := setupListener()
	if err != nil {
		log.Fatalf("[*] Error setting up listener: %v", err)
	}

	startListener(listener)
}
