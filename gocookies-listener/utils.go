package main

import (
	"fmt"
	"log"
	"net"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		log.Printf("[*] Recovered from panic: %v", r)
	}
}

func checkPortInUse(port int) bool {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return true
	}

	err = listener.Close()
	if err != nil {
		return true
	}
	return false
}
