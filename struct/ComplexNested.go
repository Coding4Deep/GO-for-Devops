package main

import "fmt"

// struct here
type students struct {
	name    string
	address string
	salary  int
}

type teacher struct {
	number int
	info   students
}

type school struct {
	studentinfo students
	teacherinfo teacher
}

func main() {

	// struct in slice
	student1 := []students{
		{name: "john", address: "westworld", salary: 2000},
	}

	fmt.Println(student1)
	fmt.Println(student1[0].name)

	// slice here
	nameofstudent := []string{"deepak", "rohit", "sagar"}
	addressofstudent := []string{"delhi", "moradabd", "jaipur"}

	fmt.Println(nameofstudent)
	fmt.Println(addressofstudent)

	// Slice of teacher structs
	teachers := []teacher{
		{number: 1, info: students{name: "Chandra-Sir", address: "Rajasthan", salary: 500000}},
	}
	fmt.Println(teachers)
	fmt.Println(teachers[0].info)
	fmt.Println("teacher name-", teachers[0].info.name)
	fmt.Println(teachers[0].info.salary)
	fmt.Println("#########################################################")

	school1 := []school{
		{studentinfo: students{name: "sweety", address: "mbd", salary: 100}, teacherinfo: teacher{number: 1, info: students{name: "newname", address: "delhi", salary: 2000}}},
	}
	fmt.Println("school1 (studentinfo)--->", school1[0].studentinfo)
	fmt.Println("school1 (studentinfo)--->", school1[0].studentinfo.name)
	fmt.Println("school1 (studentinfo)--->", school1[0].studentinfo.address)
	fmt.Println("school1 (studentinfo)--->", school1[0].studentinfo.salary)
	fmt.Println("#########################################################")
	fmt.Println("school1 (teacherinfo)--->", school1[0].teacherinfo)
	fmt.Println("school1 (teacherinfo)--->", school1[0].teacherinfo.number)
	fmt.Println("school1 (teacherinfo)--->", school1[0].teacherinfo.
		info)
	fmt.Println("school1 (teacherinfo)--->", school1[0].teacherinfo.
		info.name)
	fmt.Println("school1 (teacherinfo)--->", school1[0].teacherinfo.
		info.address)
	fmt.Println("school1 (teacherinfo)--->", school1[0].teacherinfo.
		info.salary)

}
