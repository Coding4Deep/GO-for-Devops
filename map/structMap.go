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
	}
	fmt.Println(studentinfo["student1"])
	fmt.Println(studentinfo["student1"].name)
	fmt.Println(studentinfo["student1"].section)
	fmt.Println(studentinfo["student1"].marks)

}
