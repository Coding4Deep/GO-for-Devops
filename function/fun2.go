package main

import "fmt"

func main() {

	// FUNCTION AS VARIALE

	substract := func(x, y int) int {
		return x - y
	}
	fmt.Println(substract(10, 8))

}
