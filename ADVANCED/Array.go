package main

import "fmt"

func main() {

	//DECLARATION WITHOUT INITIALLIZATION
	var array1 [5]int
	array1[0] = 10
	fmt.Println(array1[0])
	fmt.Println("#################################################")

	//DECLARATION WITH INITIALLIZATION
	var array2 = [4]int{10, 20, 30, 40}
	fmt.Println(array2)
	fmt.Println("#################################################")

	//short hand array
	array3 := [4]int{10, 20, 20, 30}
	fmt.Println(array3)
	fmt.Println("#################################################")

	//using ... operator
	array4 := [...]int{10, 10, 10}
	fmt.Println(array4)
	fmt.Println("#################################################")

	//array of array
	array5 := [4][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 18, 19},
		{10, 11, 12},
	}
	fmt.Println(array5)
	fmt.Println(array5[3][2])
	fmt.Println("#################################################")

	//pointer to array
	var array6 [4]int
	array6[0] = 1
	array6[1] = 2
	array6[2] = 3
	array6[3] = 4
	var ptr *[4]int = &array6
	fmt.Println(array6)
	fmt.Println(ptr)
	fmt.Println(&array6)
	fmt.Println(ptr[1])
	fmt.Println(&ptr[1])
	fmt.Println(&array6[1])
	fmt.Println((*ptr))
	fmt.Println((&ptr))
	fmt.Println("#################################################")

	//array of struct
	type arr struct {
		name   string
		rollno int
	}
	array7 := [2]arr{
		{"deepak", 01},
		{"sagar", 02},
	}
	fmt.Println(array7)
	fmt.Println(array7[0])
	fmt.Println(array7[0].name)
	fmt.Println(array7[0].rollno)
	fmt.Println("#################################################")

	//length of array
	fmt.Printf("%d", len(array7))
	fmt.Printf("%d", len(array6))
	fmt.Printf("%d", len(array5))
	fmt.Printf("%d", len(array4))
	fmt.Printf("%d", len(array3))
	fmt.Printf("%d", len(array2))
	fmt.Printf("\n")
	fmt.Println("#################################################")
	//taking value
	/*
	   var array8 [4]int
	   fmt.Scanf("%d",&(array8[0]))
	   fmt.Scanf("%d",&(array8[1]))
	   fmt.Scanf("%d",&(array8[2]))
	   fmt.Scanf("%d",&(array8[3]))

	   fmt.Println("Entered array:", array8)
	   fmt.Scan(&(array8[0]))
	   fmt.Scan(&(array8[1]))
	   fmt.Scan(&(array8[2]))
	   fmt.Scan(&(array8[3]))
	   fmt.Println("Entered array:", array8)
	   var name string
	   fmt.Scan(&name)
	   fmt.Println(name)
	   fmt.Scanln(&name)
	   fmt.Println(name)*/
	fmt.Println("#################################################")

}
