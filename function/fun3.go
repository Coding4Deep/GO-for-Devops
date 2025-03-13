package main

import "fmt"

func calculator(a int, b int, op func(int, int) int) int {
	return op(a, b)

}

func main() {

	// FUNCTION AS A ARGUMENTS
	add := func(x, y int) int {
		return x + y
	}
	substract := func(x, y int) int {
		return x - y
	}

	fmt.Println(calculator(1, 2, add))
	fmt.Println(calculator(1, 2, substract))

}
