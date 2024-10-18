// slices
package main

import (
	
	"fmt"
	
)

func main() {
	// fmt.Println("Hello,Abhi")
	// slices
	// declartion ,taking input,printing ,pass by reference
	var n int;
	fmt.Printf("enter the value of n:")
	fmt.Scan(&n)
	numbers :=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Printf("enter the value of numbers[%d]: ",i)
		fmt.Scan(&numbers[i])
	}
	for i:=0;i<n;i++{
		fmt.Printf("the value of %d is %d:",i,numbers[i])
		fmt.Printf("\n")
	}
	fmt.Printf("\n");
	// maps
	mp:=make(map[int]int)
	for i:=0;i<n;i++{
		mp[numbers[i]]++;
	}
	for key,value :=range mp{
		fmt.Printf("key:%d value:%d\n",key,value)
	}
}