package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error is : ", err)
		return
	}

	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)
	fmt.Println(res.ProtoMajor)
	fmt.Println(res.ProtoMinor)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
