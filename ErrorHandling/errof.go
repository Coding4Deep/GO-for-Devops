package main

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("this cant be done as b value is zero ")
	}
	return a / b, nil
}

func main() {

	/*fmt.Println("Hello, this is for error handling")
	fmt.Println("##############################################")
	fmt.Println("##############################################")
	x, err := Divide(10, 0)
	fmt.Println(x, err) */

	fmt.Println("Hello, this is for error handling")
	fmt.Println("##############################################")
	fmt.Println("##############################################")
	x, _ := Divide(10, 0)
	fmt.Println(x)
}
