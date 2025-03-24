package main

import (
	"fmt"
	"net/url"
)

func main() {
	raw_url := "https://example.com:8080/path/to/query?name=John&age=30#fragment"
	parsedURL, err := url.Parse(raw_url)
	if err != nil {
		fmt.Println("getting an error : ", err)
		return
	}
	fmt.Println("playing with Parsed url: ")

	fmt.Println("Scheme:", parsedURL.Scheme)     // https
	fmt.Println("Host:", parsedURL.Host)         // example.com:8080
	fmt.Println("Path:", parsedURL.Path)         // /path
	fmt.Println("Query:", parsedURL.RawQuery)    // name=John&age=30
	fmt.Println("Fragment:", parsedURL.Fragment) // fragment
	fmt.Println("Port:", parsedURL.Port())       // 8080
	fmt.Println("########################")

	// Query Parameters
	queryParams := parsedURL.Query()

	fmt.Println("Name:", queryParams.Get("name")) // John
	fmt.Println("Age:", queryParams.Get("age"))   // 30

}
