package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("getting error", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("status is :", res.Status)
	fmt.Printf(" type of ResponseData is : %T", res)
	ResponseData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error is :", err)
		return
	}

	fmt.Println(string(ResponseData))

}
