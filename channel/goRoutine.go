package main

import (
	"fmt"
	"time"
)

// Function that prints numbers
func printNumbers() {
	for i := 1; i <= 9; i++ {
		fmt.Println(i)
		time.Sleep(50 * time.Second) // Simulate work
	}
}

func main() {
	go printNumbers() // Start a goroutine

	// Main function continues execution
	fmt.Println("Main function is running...")

	time.Sleep(100 * time.Second) // Wait for goroutine to finish
	fmt.Println("Main function ends")
}
