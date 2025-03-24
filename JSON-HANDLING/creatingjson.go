package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	p1 := person{
		Name:  "deepak",
		Age:   24,
		Email: "sagardeepak2002@gmail.com",
	}

	//ENCODING JSON
	JSON_DATA, err := json.MarshalIndent(p1, "", "  ") //for pretty print
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(JSON_DATA))
	fmt.Println("########################")

	//DECODING JSON

}
