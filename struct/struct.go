package main

import "fmt"

type students struct {
	name   string
	rollno int
	marks  string
}

type persons struct {
	name string
	age  int
	city string
}

func main() {
	student1 := students{"deepak", 101, "A"}

	student2 := students{
		name:   "sagar",
		rollno: 101,
		marks:  "A",
	}

	student3 := students{name: "rohit", rollno: 103, marks: "B"}

	person1 := persons{"deepak sagar", 25, "pune"}

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student3)
	fmt.Println("#############################################")
	fmt.Println("Name:", student1.name)
	fmt.Println("Rollno:", student1.rollno)
	fmt.Println("Marks:", student1.marks)
	fmt.Println("#############################################")
	fmt.Println("name:", person1.name)
	fmt.Println("age:", person1.age)
	fmt.Println("city:", person1.city)
	fmt.Println(person1)

}
