package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Method associated with the Person struct
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	p := Person{Name: "Deepak", Age: 25}
	p.Greet() // Calling the method
}
