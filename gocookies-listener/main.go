package main

import (
	"context"
	"github.com/simonfalke-01/gocookies/listener/redis"
	"log"
)

func main() {
	// gets argv[1] from cli as the port to run the redis container on
	// gets argv[2] from cli as the port to run the listener on
	redisPort, listenerPort, useExisting := getFlags()

	// start redis container
	if !useExisting {
		if checkPortInUse(redisPort) {
			log.Fatalf("[*] Redis port %v is already in use", redisPort)
		}

		_, err := redis.CreateRedisContainer(redisPort)
		if err != nil {
			log.Fatalf("[*] Error creating redis container: %v", err)
		}
	}

	if checkPortInUse(redisPort) {
		log.Fatalf("[*] Listener port %v is already in use", redisPort)
	}

	listener, err := setupListener(listenerPort)

	if err != nil {
		log.Fatalf("[*] Error setting up listener: %v", err)
	}

	c := redis.NewClient(context.Background(), "localhost", redisPort, "", 0)

	startListener(listener, c)
}
