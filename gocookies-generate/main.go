package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	host := flag.String("h", "", "host")
	port := flag.Int("p", 0, "port")
	dir := flag.String("d", "", "directory")
	verbose := flag.Bool("v", false, "verbose")

	flag.Parse()

	if *host == "" || *port == 0 || *dir == "" {
		fmt.Println("All flags (-h, -p, -d) are required")
		os.Exit(1)
	}

	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		fmt.Println("Directory does not exist:", *dir)
		os.Exit(1)
	}

	fmt.Printf("Building gocookies with host=%s, port=%d...\n", *host, *port)

	buildCmd := exec.Command(
		"go", "build",
		"-ldflags",
		fmt.Sprintf("-X main.host=%s -X main.port=%d -X main.verbose=%v", *host, *port, func(verbose *bool) string {
			if *verbose {
				return "true"
			}
			return "false"
		}(verbose)),
		"-o", filepath.Join(*dir, "output"))

	buildCmd.Dir = *dir
	buildOut, err := buildCmd.CombinedOutput()

	if err != nil {
		fmt.Println("Build failed:", err)
		fmt.Println(string(buildOut))
		os.Exit(1)
	}

	// Move the compiled program to the current directory
	if err := os.Rename(filepath.Join(*dir, "output"), "./output"); err != nil {
		fmt.Println("Failed to move compiled program:", err)
		os.Exit(1)
	}

	fmt.Println("gocookies built and moved to current directory.")
}
