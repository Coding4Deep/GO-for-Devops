package main

import "fmt"

type students struct {
	name    string
	section string
	marks   int
}

func main() {
	studentinfo := map[string]students{
		"student1": {"deepak", "B", 23},
		"student2": {"sagar", "C", 24},
		"student3": {"rohan", "D", 25},
		"student4": {"mohan", "E", 26},
		"student5": {"sohan", "F", 27},
		"student6": {"kohan", "G", 28},
	}
	fmt.Println(studentinfo["student1"])
	fmt.Println(studentinfo["student1"].name)
	fmt.Println(studentinfo["student1"].section)
	fmt.Println(studentinfo["student1"].marks)
	fmt.Println("            ")
	fmt.Println("#################################")
	fmt.Println("#################################")
	fmt.Println("            ")

}

// Output:
