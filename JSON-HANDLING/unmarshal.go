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
	jsonString := `{
       "name":"deepak",
       "age":24,
       "email":"sagardeepak2002@gmail.com"
   }`
	p2 := person{}
	err := json.Unmarshal([]byte(jsonString), &p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)

}
