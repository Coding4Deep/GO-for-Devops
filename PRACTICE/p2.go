package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type person struct {
	UserID   int    `json:"userid"`
	Id       int    `json:"Id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

func httpGET() {
	GETres, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("getting GET error : ", err)
		return
	}
	defer GETres.Body.Close()

	GET_DATA, err := io.ReadAll(GETres.Body)
	if err != nil {
		fmt.Println("getting GET error : ", err)
		return
	}
	fmt.Println("#########################")
	fmt.Println(string(GET_DATA))
	fmt.Println("#########################")

}

func jsonPARSE() string {
	p1 := person{
		UserID:   01,
		Id:       101,
		Title:    "this is for everything",
		Complete: false,
	}
	json_DATA, err := json.MarshalIndent(p1, "", " ")
	if err != nil {
		fmt.Println("getting error : ", err)
		return
	}
	fmt.Println("#########################")
	fmt.Println("this is parsed data : ", string(json_DATA))
	fmt.Println("#########################")

	return string(json_DATA)

}

func httpPOST(s string) {
	myURL := ("https://jsonplaceholder.typicode.com/todos/")
	jsonReader := strings.NewReader(s)
	POST_RES, err := http.Post(myURL, "application/json", jsonReader)

	if err != nil {
		fmt.Println(err)
	}
	defer POST_RES.Body.CLose()

	fmt.Println(POST_RES.Status)
	fmt.Println(string(POST_RES))

}

func main() {
	httpGET()
	jsonString := jsonPARSE()
	httpPOST(jsonString)

}
