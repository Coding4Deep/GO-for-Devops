package main

import "fmt"

func main() {

	var s1 string = "hello"
	s2 := "world"
	s3 := "hello world \nhow are you?"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	s4 := `hello world \nhow are you?`
	fmt.Println(s4)
}
