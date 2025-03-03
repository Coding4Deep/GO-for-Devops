package main

import "fmt"

func main() {

	s1 := "Hello"
	s2 := "World"
	s3 := s1 + " " + s2 // Concatenation using `+`
	fmt.Println(s3)     // Output: Hello World
	str := "GoLang"
	fmt.Println(len(str)) // Output: 6 (number of bytes, not characters)
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(len(s3))

	fmt.Println(string(str[0])) // Output: G
	fmt.Println(string(str[3])) // Output: a
	fmt.Println(string(str[5])) // Output: g
	fmt.Println(string(s1[0]))  // Output: H
	fmt.Println(string(s2[3]))  // Output: l

	fmt.Println(s1 == s2) // false
	fmt.Println(s1 == s3) // true
	s7 := "Hello"
	fmt.Println(s1 == s7) // true

}
