package main

import "fmt"

func main() {
	slice := []string{"deepak", "rahul", "rohit", "mohit", "sahil"}
	fmt.Println(slice)

	newslice := append(slice[0:2], "newname") // append function is used to add new element in slice
	fmt.Println(newslice)
	fmt.Println(len(newslice))
}
