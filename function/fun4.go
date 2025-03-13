package main

import "fmt"

func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {

	// RETURNING A FUNCTION FROM ANOTHER FUNCTION

	timesTwo := getMultiplier(2)
	fmt.Println(timesTwo(5))

	timesFive := getMultiplier(5)
	fmt.Println(timesFive(5))

}
