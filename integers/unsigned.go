package main

import "fmt"

func main() {
	var a uint8 = 255
	var b uint16 = 65535
	var c uint32 = 4294967295
	var d uint64 = 18446744073709551615

	fmt.Println(a, b, c, d)
}
