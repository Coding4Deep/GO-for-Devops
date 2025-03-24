package main

import (
	"fmt"
	"os"
	//   "bufio"
	// "io"
)

func main() {
	fmt.Println("Creating a File")
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(file.Name(), "created")

	// fmt.Println("Writing to a File")
	_, error := file.WriteString("Hello, this is for error handling")
	if error != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Writing to a File")

}
