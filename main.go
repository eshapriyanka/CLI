package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Define a flag: user can type --url to check a specific site
	urlPtr := flag.String("url", "https://hennge.com", "URL to check")
	flag.Parse()

	fmt.Printf("Checking status for: %s ...\n", *urlPtr)

	// Set a timeout of 2 seconds
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	// Make a GET request
	resp, err := client.Get(*urlPtr)
	if err != nil {
		fmt.Printf("[DOWN] Could not reach %s\nError: %v\n", *urlPtr, err)
		os.Exit(1) // Exit with error code (Unix style)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode == 200 {
		fmt.Printf("[UP] Website is online! (Status: %d)\n", resp.StatusCode)
		os.Exit(0) // Exit with success code
	} else {
		fmt.Printf("[WARNING] Website returned status: %d\n", resp.StatusCode)
		os.Exit(1)
	}
}