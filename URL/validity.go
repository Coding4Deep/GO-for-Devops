package main

import (
	"fmt"
	"net/url"
)

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
func main() {
	fmt.Println(isValidURL("https://golang.org")) // true
	fmt.Println(isValidURL("invalid-url"))        // false
}
