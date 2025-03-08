package main

import "fmt"

func newValue(b *int) {
	*b = *b + 10
	fmt.Println(b)

}

func main() {
	a := 50
	fmt.Println("before", a)
	newValue(&a)
	fmt.Println("after", a)
	a = 70
	fmt.Println("after", a)
	newValue(&a)
	fmt.Println("after", a)
	c := 200
	fmt.Println(c)
	newValue(&c)
	fmt.Println(c)

}
