package main

import (
	"net/url"
)

func main() {
	url := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "Path/to/API",
		RawQuery: url.Values{
			"name": []string{"John"},
			"age":  []string{"30"},
		}.Encode(),
	}
	println(url.String())
}
