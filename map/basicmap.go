package main

import "fmt"

func main() {
	marks := map[string]int{
		"rohan": 30,
		"mohan": 31,
	}

	fmt.Println(marks["rohan"])
	fmt.Println(marks["mohan"])
	fmt.Println(marks)

}
