package main

import "fmt"

type employee struct {
	name   string
	salary int
}

func updateemployee(emp *employee) {
	emp.salary += 100
}

func addname(newemp *employee) {
	*newemp = employee{"sagar", 2000}
}

func main() {
	emp1 := employee{"deepak", 100}
	fmt.Println(emp1)
	//AFTER UPDATE
	updateemployee(&emp1)
	fmt.Println(emp1)

	//ADDING NEW EMPLOYEE USING POINTER

	var em2 employee
	addname(&em2)
	fmt.Println(em2)

}
