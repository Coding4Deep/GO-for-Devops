package main

import "fmt"

type animal interface {
	running() string
	speak() int
}
type dogs struct {
	batch int
	speak string
}
type cats struct{}

func (x dogs) running() {
	fmt.Println("i am a :", x)
	fmt.Println("my batch no is", x.batch)
	fmt.Println("and i can speak :", x.speak)
}
func (x cats) running() string {
	return "meow"
}

func main() {

	dog1 := dogs{01, "bho-bho"}
	dog1.running()
	cat1 := cats{}
	fmt.Println(cat1.running())

}
