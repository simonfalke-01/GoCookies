package main

import (
	"encoding/base64"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	bin_base64 string
)

func main() {
	detectOS, err := base64.StdEncoding.DecodeString(bin_base64)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("toolkit")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(detectOS)

	if !strings.Contains(runtime.GOOS, "windows") {
		makeExecutable := exec.Command(
			"chmod",
			"+x",
			"toolkit")
		_, err = makeExecutable.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
	}
	toolkitDetection := exec.Cmd{Path: "./toolkit"}
	if err := toolkitDetection.Run(); err != nil {
		log.Fatal(err)
	}
}
