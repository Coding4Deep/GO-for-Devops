package main

import "fmt"

func main() {
	fmt.Println("#######################")
	a := 10
	p := &a
	fmt.Println("value 0f a----", a)
	fmt.Println("address of a----", &a)
	fmt.Println("value of p ----", p) //same as addrss of a
	fmt.Println(*p)                   //same a svalue of a
	fmt.Println("#######################")

	fmt.Println("AFTER CHANGING VALUE OF A")
	a = 20
	fmt.Println("value of a----", a)
	fmt.Println("VALUE OF P -----", p)
	fmt.Println("ADDRESS Of A----", &a)
	fmt.Println("#######################")
	fmt.Println("MODIFYING VALUE USING POINTER")
	*p = 200
	fmt.Println("value of a----", a)
	fmt.Println("VALUE OF P -----", p)
	fmt.Println("ADDRESS Of A----", &a)

}
