package main

import "fmt"

func main() {
	var x int8 = -128
	var y int16 = 32767
	var z int32 = 2147483647
	var w int64 = 9223372036854775807

	fmt.Println(x, y, z, w)
}
