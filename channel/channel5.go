package main

import (
	"fmt"
	"time"
)

func newfunction() {
	fmt.Println("hello world")
}

func ndfunc(channel2 chan int) {
	channel2 <- 1
	channel2 <- 2
	channel2 <- 3
	channel2 <- 4
	channel2 <- 5
}

func main() {
	go newfunction()
	time.Sleep(1 * time.Second)
	fmt.Println("main func end")
	channelname := make(chan int)
	channel2 := make(chan int)
	//go ndfun()

	go func() {
		channelname <- 1
		channelname <- 2
		channelname <- 3
		channelname <- 4
		channelname <- 5
	}()
	fmt.Println(<-channelname)
	fmt.Println(<-channelname)
	fmt.Println(<-channelname)
	fmt.Println(<-channelname)
	fmt.Println(<-channelname)
	fmt.Println("###############################")

	go ndfunc(channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)

}
