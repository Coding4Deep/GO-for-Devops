package main

import "fmt"

// The zero value is the default value of a variable when it is declared

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)
	fmt.Printf("%d\n%f\n%t\n%s\n", i, f, b, s)
	fmt.Printf("%.2f\n", f) // Prints float with 2 decimal places

	/*
	   %d	Integer (int)
	   %f	Floating-point (float64)
	   %t	Boolean (bool)
	   %s	String (string)
	   \n	Newline (moves to next line)	(Newline effect)
	*/

}
