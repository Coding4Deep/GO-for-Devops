package main

import "fmt"

type students struct {
	name   string
	rollno int
}

func (x students) sayHello() {
	fmt.Println("my name is : ", x.name)
	fmt.Println("my roll number is : ", x.rollno)

}

func main() {
	student1 := students{
		name:   "deepak",
		rollno: 1,
	}
	student1.sayHello()
	student2 := students{
		name:   "sagar",
		rollno: 12,
	}
	student2.sayHello()
}
