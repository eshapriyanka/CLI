package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	urlPtr := flag.String("url", "", "The URL to check")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Println("Error: Please provide a URL using --url")
		os.Exit(1)
	}

	fmt.Printf("Checking %s ...\n", *urlPtr)
	client := http.Client{Timeout: 2 * time.Second}

	resp, err := client.Get(*urlPtr)
	if err != nil {
		fmt.Printf("[DOWN] Error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("[UP] Status: %d OK\n", resp.StatusCode)
		os.Exit(0)
	} else {
		fmt.Printf("[WARNING] Status: %d\n", resp.StatusCode)
		os.Exit(1)
	}
}