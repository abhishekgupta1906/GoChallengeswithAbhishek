package main

import (
	
	"fmt"
	
)

func main() {
	 arr :=[]int{1,2,4}
	
	 fmt.Printf("len is %d\n and capacity is %d\n", len(arr),cap(arr))

	 arr=append(arr,7)
	 fmt.Printf("len is %d\n and capacity is %d\n", len(arr),cap(arr))
}