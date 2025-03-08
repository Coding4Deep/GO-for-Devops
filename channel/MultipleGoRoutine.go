package main

import (
	"fmt"
	"time"
)

func task1() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Task 1 -", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func task2() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Task 2 -", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go task1() // Run as goroutine
	go task2() // Run as goroutine

	time.Sleep(2 * time.Second) // Give time for goroutines to complete
	fmt.Println("Main function exits")
}
