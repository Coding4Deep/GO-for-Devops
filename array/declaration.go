package main

import "fmt"

func newfunction() {
	slice := []int{10, 222, 34, 41, 5}
	fmt.Println(slice)
	newslice := slice[1:3]
	fmt.Println(newslice)
}

func main() {
	students_name := [5]string{"deepak", "rahul", "rohit", "mohit", "sahil"}
	students_rollno := [5]int{1, 2, 3, 4, 5}

	fmt.Println(students_name)
	fmt.Println(students_rollno)

	no := "#############################################"
	fmt.Println(no)
	newfunction()
	fmt.Println(no)
	numbers := [6]int{21, 212, 24, 94, 303, 33}
	fmt.Println(numbers)
	fmt.Println(students_name[2])

}
