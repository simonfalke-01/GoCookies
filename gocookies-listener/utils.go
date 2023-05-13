package main

import (
	"log"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		log.Printf("[*] Recovered from panic: %v", r)
	}
}
