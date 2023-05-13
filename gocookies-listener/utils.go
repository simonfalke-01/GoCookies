package main

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"time"
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

func extractHost(str string) (string, error) {
	re := regexp.MustCompile(`^([a-zA-Z0-9.\[\]:]+):\d+$`)
	match := re.FindStringSubmatch(str)
	if len(match) < 2 {
		return "", fmt.Errorf("invalid input string")
	}
	return match[1], nil
}

func getCurrentTime() string {
	return time.Now().Format("2006-01-02T15:04:05Z07:00")
}
