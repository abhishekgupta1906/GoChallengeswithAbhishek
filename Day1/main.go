package main

import (
	"fmt"
	"reflect"
)

// variable is golang
func main() {
	fmt.Println("Hello,bhai")
	const name string = "Abhi"
	fmt.Println("My name is:", name)
	fmt.Println("type of name is:", reflect.TypeOf(name))

	// for exporting variables we intitlize the name of varibale with capital letter
	var Age int = 4
	fmt.Printf("My name is: %s and age is %d\n", name, Age)

  // one more way of variable declaration
  height :=4
  height=3
  fmt.Println(height)
  fmt.Println("type of height is:", reflect.TypeOf(height))


}
