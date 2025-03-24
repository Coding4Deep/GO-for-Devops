package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"

	// Check if the file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
		return
	}

	// Open the file in append mode (write)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write to file
	_, writeErr := file.WriteString("\nHello, this is for error handling from another calling function.")
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return
	}

	fmt.Println("Successfully written to file.")
}
