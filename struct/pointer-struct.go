package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	p := Person{Name: "Deepak", Age: 23}
	p.HaveBirthday()               // Increments Age
	fmt.Println("New Age:", p.Age) // 26
}
