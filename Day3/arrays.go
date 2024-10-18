package main

import (
	"fmt"
)

// array declaration
func append_at_index(arr *[]int, index int, value int) {
	// already by refernce gya h
    (*arr)[0] =4
	*arr = append((*arr)[:index], append([]int{value}, (*arr)[index:]...)...)

	// fmt.Println(arr);
}
func delete_last(arr *[]int) {
	if len(*arr) > 0 {
        *arr = (*arr)[:len(*arr)-1]
    }
	// fmt.Println(arr);
}
func main() {
	fmt.Println("Hello,Abhi")
	// 1st method
	var numbers = []int{1, 2, 3, 4}
	fmt.Println(numbers)
	// 2nd method
	// keys :=[]string{"abc","def","tjk"};
	keys := []string{}

	fmt.Println(keys)

	// appending in arrays
	numbers = append(numbers, 9)
	fmt.Println(numbers)
	
	// appending at any index
	append_at_index(&numbers, 2, 15)
	fmt.Println(numbers)

	// deleting the last number
	delete_last(&numbers)
	fmt.Println(numbers)

}
