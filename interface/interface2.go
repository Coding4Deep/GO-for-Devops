package main

import "fmt"

func describe(i interface{}) {
	if v, ok := i.(string); ok { // Type assertion for string
		fmt.Println("String value:", v)
	} else if v, ok := i.(int); ok { // Type assertion for int
		fmt.Println("Integer value:", v)
	} else {
		fmt.Println("Not a string, neither an integer")
	}
}

func main() {
	describe("Hello") // Output: String value: Hello
	describe(100)     // Output: Integer value: 100
	describe(3.14)    // Output: Not a string, neither an integer
}
