package main

import (
	"fmt"
	"time"
)

func main() {
	// Declare and initialize channels
	channel1 := make(chan int, 3) // Buffered channel with capacity 3
	channel2 := make(chan int)

	// Goroutine to send a value into channel1
	go func() {
		channel1 <- 100
	}()

	// Receiving from channel1 (this prevents deadlock)
	variable1 := <-channel1
	fmt.Println("Received from channel1:", variable1)

	// Sending multiple values (buffered channel avoids deadlock)
	channel1 <- 101
	channel1 <- 102
	channel1 <- 103

	// Receiving the sent values
	fmt.Println("Received:", <-channel1)
	fmt.Println("Received:", <-channel1)
	fmt.Println("Received:", <-channel1)

	// Goroutine to prevent blocking send on channel2
	go func() {
		channel2 <- 200
	}()

	fmt.Println("Received from channel2:", <-channel2)

	// Wait to allow goroutines to complete (optional)
	time.Sleep(1 * time.Second)
}
