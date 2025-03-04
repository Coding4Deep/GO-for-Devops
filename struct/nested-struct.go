package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Person struct {
	Name    string
	Age     int
	Address Address // Nested struct
}

func main() {
	p := Person{Name: "Deepak", Age: 25, Address: Address{City: "Delhi", State: "India"}}
	fmt.Println(p.Name, "lives in", p.Address.City)
}
