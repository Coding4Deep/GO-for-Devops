package main

import "fmt"

//struct 1
type students struct {
	name   string
	rollno int
	grade  string
}

//struct 2
type person struct {
	name string
	age  int
	city string
}

func main() {
	//map
	rn := map[string]int{
		"deepak":    18,
		"sagar":     19,
		"priya":     20,
		"priyanshu": 21,
	}
	//playing with map
	fmt.Println(rn)
	fmt.Println(rn["deepak"])
	fmt.Println(rn["priya"])
	fmt.Println("#################################")
	fmt.Println("#################################")
	fmt.Println("#################################")

	//playing with struct + map together

	info := map[string]person{
		"person1": {name: "deepak", age: 222, city: "morada"},
		"person2": {name: "sagar", age: 333, city: "mabs"},
	}
	fmt.Println(info["person1"].name)
	fmt.Println(info["person2"])
	fmt.Println(info["person2"].city)
	fmt.Println(info["person2"].age)
	fmt.Println(info["person1"].city)
	fmt.Println(info)
	fmt.Println("###################################")
	fmt.Println("###################################")
	fmt.Println("###################################")

	//struct only
	student1 := students{"deepak", 001, "A"}
	fmt.Println("this is student 1 :", student1)
	student2 := students{
		name:   "deepaksagar",
		rollno: 23,
		grade:  "B",
	}
	fmt.Println("this is student 2 :", student2.name)
	fmt.Println(student2.rollno)
	fmt.Println(student2.grade)
	fmt.Println(student2)
	fmt.Println("###################################")
	fmt.Println("###################################")
	fmt.Println("###################################")

	//struct in slice

	student3 := []students{
		{name: "rohit", rollno: 12, grade: "D"},
		{name: "sohit", rollno: 13, grade: "e"},
		{name: "mohit", rollno: 14, grade: "f"},
		{name: "pohit", rollno: 15, grade: "g"},
	}

	fmt.Println(student3[0].name)
	fmt.Println(student3[0])
	fmt.Println(student3[0].rollno)
	fmt.Println(student3[0], student3[1])

}
