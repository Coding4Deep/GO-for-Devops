package main

import (
	"fmt"
	"time"
)

type employee struct {
	name   string
	salary int
}

// Function to send employee data to the channel
func getValues(a int, b chan employee) {
	b <- employee{fmt.Sprintf("Employee-%d", a), 1000}
	//time.Sleep(2 * time.Second)
}

func main() {
	channel := make(chan employee)

	for i := 1; i <= 5; i++ {
		go getValues(i, channel)
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(<-channel)
		time.Sleep(1 * time.Second)
	}
}
