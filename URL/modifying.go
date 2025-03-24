package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://example.com/api")
	// Modify query parameters
	queryParams := u.Query()
	queryParams.Set("name", "DEEPAK")
	queryParams.Set("age", "22")
	u.RawQuery = queryParams.Encode()        // Encode modified query
	fmt.Println("Modified URL:", u.String()) // https://example.com/api?id=123&page=2
}
