package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	bin := flag.String("b", "", "Path to gocookies binary")
	dir := flag.String("d", "", "path to gocookies-wrapper directory")
	if *bin == "" || *dir == "" {
		fmt.Println("not enough arguments provided.")
	}
	bin_contents, err := os.ReadFile(*bin)
	if err != nil {
		log.Fatal(err)
	}
	bin_base64 := base64.StdEncoding.EncodeToString([]byte(bin_contents))
	buildCmd := exec.Command(
		"go",
		"build",
		"-ldflags",
		fmt.Sprintf("-X main.bin_base64=%s", bin_base64),
	)

	buildCmd.Dir = *dir
	_, err = buildCmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}
