package main

import "fmt"

// this program include evrything ARRAY,SLICE
//STRUCT,MAP,POINTER,CHANNEL,GOROUTINES,

type struct1 struct {
	name   string
	rollno int
	marks  float32
}
type struct2 struct {
	students struct1
}

func main() {

	fmt.Println("############################")
	array1 := [4]string{"deepak", "sagar", "deepaksagar", "priyanshu"}
	array2 := [...]int{1, 2, 3, 4}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array1[0], array2[0])
	fmt.Println("############################")

	slice1 := []string{"deepak", "sagar", "deepaksagar", "priyanshu"}
	fmt.Println(slice1)
	fmt.Println("############################")

	//STRUCT IN ARRAY

	array3 := [2]struct1{
		{name: "deepak", rollno: 1, marks: 17.6},
		{name: "sagar", rollno: 2, marks: 18.6},
	}
	// Print the array
	fmt.Println("array struct start from here")
	fmt.Println(array3)
	fmt.Println(array3[0].name)
	fmt.Println(array3[0].rollno)
	fmt.Println(array3[0].marks)
	fmt.Println(array3[1].name)
	fmt.Println(array3[1].rollno)
	fmt.Println(array3[1].marks)

	//STRUCT IN SLICE
	slice2 := []struct1{
		{name: "deepak", rollno: 1, marks: 17.6},
		{name: "sagar", rollno: 2, marks: 18.6},
	}
	fmt.Println("############################")
	fmt.Println("slice struct start from here")
	fmt.Println(slice2)
	fmt.Println(slice2[0].name)
	fmt.Println(slice2[1].name)

	//MAP
	fmt.Println("############################")
	fmt.Println("############################")
	newmap := map[struct1]struct2{
		{name: "deepak", rollno: 1, marks: 2.7}: {students: struct1{name: "sagar", rollno: 2, marks: 3.7}},
	}
	fmt.Println(newmap)
	key := struct1{name: "deepak", rollno: 1, marks: 2.7}
	fmt.Println(newmap[key].students.name)

}
