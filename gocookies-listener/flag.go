package main

import (
	"flag"
	"fmt"
	"os"
)

func getFlags() (int, int, bool) {
	redisPort := flag.Int("r", 0, "Redis container port")
	listenerPort := flag.Int("p", 0, "Listener port")
	existingRedis := flag.Bool("e", false, "Existing Redis db flag")

	flag.Parse()

	// Check if required flags are provided
	if *redisPort == 0 || *listenerPort == 0 {
		fmt.Println("Usage: go run main.go -r <redis port> -p <listener port> [-e]")
		os.Exit(1)
	}

	useExisting := false
	if *existingRedis {
		useExisting = true
	}

	return *redisPort, *listenerPort, useExisting
}
