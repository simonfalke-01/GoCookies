package main

import (
	"log"
	"net"
)

func setupListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func startListener(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer recoverFromPanic()

			handleConnection(conn)
		}()
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("[*] Error closing connection: %v", err)
		}
	}(conn)

	data := readAll(conn)
	log.Printf("[*] Received data: %s", string(data))
}

func readAll(conn net.Conn) []byte {
	var data []byte
	var buf [1024]byte

	for {
		length, err := conn.Read(buf[0:])
		if err != nil {
			log.Printf("[*] Error reading from connection: %v", err)
			break
		}

		data = append(data, buf[0:length]...)

		if length == 0 {
			break
		}
	}

	return data
}
