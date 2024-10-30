package main

import (
	
	"fmt"
	
	
)
func print_map(mp *map[int]int){
	for key,val :=range *mp{
		fmt.Printf("key %d :%d",key,val)
		fmt.Printf("\n")
	}
}
func main() {
	fmt.Println("Hello, world")
	// create a map
	arr :=[5]int{2,2,5,6,7}
	mp :=make(map[int]int)
	for i,val :=range arr{
		fmt.Printf("index is %d and value is %d",i,val)
		fmt.Printf("\n")

	}
	for i :=range arr{
		mp[arr[i]]++
	}
	for key,val :=range mp{
		fmt.Printf("key %d :%d",key,val)
		fmt.Printf("\n")
	}
	// check if a key exist or not
	_,exist :=mp[3]
	if exist{
		fmt.Println("key exist")
	}else{
		fmt.Println("key does not exist")
	}
	delete(mp,2)
	print_map(&mp)
	// map definition and intilization
	mp1 :=map[int]int{
		1:2,
		3:7,
	}
	print_map(&mp1)



}