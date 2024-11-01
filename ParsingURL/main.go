package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Hello, world")
	myurl := "https://jsonplaceholder.typicode.com/todos?id=4"
	parsedurl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(parsedurl.Scheme)
	fmt.Println(parsedurl.Host)
	fmt.Println(parsedurl.Path)
	fmt.Println(parsedurl.RawQuery)
	queryParams := parsedurl.Query()
	fmt.Println(queryParams)
	id := queryParams.Get("id")
	fmt.Println(id)
}
