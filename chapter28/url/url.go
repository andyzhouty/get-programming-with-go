package main

import (
	"fmt"
	"net/url"
)

func checkURL(URL string) {
	fmt.Printf("Parsing '%v'.\n", URL)
	result, err := url.Parse(URL)
	if err != nil {
		fmt.Printf("Error: %#v\n", err)
		if e, ok := err.(*url.Error) ; ok {
			fmt.Printf("Error Details: %v.", e)
			return
		}
	}
	fmt.Printf("URL %v is valid.\n", result)
}

func main() {
	checkURL("https://example.com")
	checkURL("https://a b.com")
}
