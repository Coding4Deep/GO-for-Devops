package main

import "fmt"

func main() {
	a := 10 // Declare a normal variable
	b := &a // Correct way to declare a pointer
	c := &b
	var d **int = &b
	fmt.Println("Value of a:", a)       // 10
	fmt.Println("Address of a:", &a)    // Address of 'a'
	fmt.Println("Value of b:", b)       // Address of 'a'
	fmt.Println("Dereferencing b:", *b) // 10 (value at address)

	fmt.Println("###############################")
	fmt.Println("Value of c:", c)                //address of b
	fmt.Println("Value of c is value of b:", *c) //value of b
	fmt.Println("Value a:", **c)                 //address of b
	fmt.Println("Value of d:", d)                //address of b
}
