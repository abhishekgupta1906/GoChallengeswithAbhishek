package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
)

func main() {
	res,err :=http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err!=nil{
		fmt.Println("Error",err)
		return
	}
	defer res.Body.Close()
	// what is this response
	// fmt.Println(res)
    data,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("Error",err)
		return
	}
	// body-bytes array h,
	fmt.Println(string(data))
}