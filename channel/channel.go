package main

import "fmt"

func main() {
	newchannel := make(chan int)
	fmt.Println(newchannel)

	go func() {
		newchannel <- 20
	}()

	x := <-newchannel
	fmt.Println(x)
}
