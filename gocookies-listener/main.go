package main

import (
	"context"
	"github.com/simonfalke-01/gocookies/listener/redis"
	"log"
	"os"
	"strconv"
)

func main() {
	// gets argv[1] from cli as the port to run the redis container on
	// gets argv[2] from cli as the port to run the listener on
	if len(os.Args) != 3 {
		log.Fatalf("[*] Usage: %s <redis port> <listener port>", os.Args[0])
	}

	redisPort := func() int {
		port, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("[*] Error converting redis port to int: %v", err)
		}
		return port
	}()
	listenerPort := func() int {
		port, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("[*] Error converting listener port to int: %v", err)
		}
		return port
	}()

	// start redis container
	_, err := redis.CreateRedisContainer(redisPort)

	if err != nil {
		log.Fatalf("[*] Error creating redis container: %v", err)
	}

	listener, err := setupListener(listenerPort)

	if err != nil {
		log.Fatalf("[*] Error setting up listener: %v", err)
	}

	c := redis.NewClient(context.Background(), "localhost", redisPort, "", 0)

	startListener(listener, c)
}
