/*
The Exercise
Write a Go program that prints a server startup message like a real backend server would.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Seth"
	port := "http://localhost:8080"
	var environment = "development"
	fmt.Println("Starting server...")
	fmt.Printf("Server name:    %s\n", name)
	fmt.Printf("Listening on:   %s\n", port)
	fmt.Printf("Environment:    %s\n", environment)
	fmt.Printf("Started at:     %s\n", time.Now().Format("2006-01-02 15:44:00"))

}
