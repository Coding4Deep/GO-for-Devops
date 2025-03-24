package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status:", res.Status)

	//READINF FROM THE RESPONSE
	//fmt.Println(res)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("erros is :", err)
		return
	}
	fmt.Println(string(data))

}
