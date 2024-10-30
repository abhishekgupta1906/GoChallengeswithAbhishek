package main

import (
	
	"fmt"
	"sort"
	
)
func twoSum(nums []int, target int) []int {
	mp1 :=make(map[int]int)
	for i,val :=range nums{
         mp1[val]=i
	}
    sort.Ints(nums)
	i := 0
	j :=len(nums)-1
	ans :=make([]int,2,5)
	for i<j{
		if nums[i]+nums[j]==target{
			ans[0]=mp1[nums[i]]
			ans[1]=mp1[nums[j]]
			break
		}else if nums[i]+nums[j]>target{
			j--
		}else{
			i++
		}


	}
	return ans	


}

func main() {
	fmt.Println("Hello, world")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums,target))
}