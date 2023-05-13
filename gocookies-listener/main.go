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

		log.Printf("[*] Redis container running on port %v", redisPort)
	} else {
		redisDbUp := redis.TestRedisConnection("localhost", redisPort)

		if !redisDbUp {
			log.Fatalf("[*] Redis db not running on port %v", redisPort)
		}
	}

	if checkPortInUse(listenerPort) {
		log.Fatalf("[*] Listener port %v is already in use", listenerPort)
	}

	listener, err := setupListener(listenerPort)

	if err != nil {
		log.Fatalf("[*] Error setting up listener: %v", err)
	}

	c := redis.NewClient(context.Background(), "localhost", redisPort, "", 0)

	log.Printf("[*] Listener running on port %v", listenerPort)

	startListener(listener, c)
}
