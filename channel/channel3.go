package main

import (
	"fmt"
	"time"
)

// ✅ printingNumbers now takes a channel
func printingNumbers(ch3 chan int) {
	for i := 0; i < 10; i++ {
		ch3 <- i // Send numbers to channel
		time.Sleep(500 * time.Millisecond)
	}
	close(ch3) // Close the channel after sending numbers
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan int) // ✅ Created ch3

	// Goroutine for ch1
	go func() {
		ch1 <- "Hello from ch1"
		//	time.Sleep(2 * time.Second)
		ch1 <- "Hello from ch1 again"
	}()

	// Goroutine for ch2
	go func() {
		//	time.Sleep(1 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	// ✅ Receiving values in main()
	fmt.Println(<-ch1) // Receives "Hello from ch1"
	fmt.Println(<-ch1) // Receives "Hello from ch1 again"
	fmt.Println(<-ch2) // Receives "Hello from ch2"

	// ✅ Start printingNumbers() in a goroutine
	go printingNumbers(ch3)

	// ✅ Read values from ch3 in a loop
	for num := range ch3 {
		fmt.Println("Received from ch3:", num)
	}

	// ✅ Prevent main() from exiting too soon
	time.Sleep(1 * time.Second)
}
