package main

import (
	"fmt"
	"sort"
)

// change the value from pass by reference

func changeValue(arr *[]int) {
	(*arr)[0] = 5
}
func printmatrix(arr *[][]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len((*arr)[i]); j++ {
			fmt.Printf("%d ", (*arr)[i][j])

		}
		fmt.Printf("\n")
	}
}
func main() {
	arr := []int{1, 2, 4}

	fmt.Printf("len is %d\n and capacity is %d\n", len(arr), cap(arr))

	arr = append(arr, 7)
	fmt.Printf("len is %d\n and capacity is %d\n", len(arr), cap(arr))
	//  or make pair method
	arr1 := make([]int, 3, 5)
	arr1[0] = 4
	for i,val:= range arr1 {
		fmt.Println("arr1 value is ",val,i)
	}
	changeValue(&arr1)
	fmt.Println(arr1)

	// 2d arr or matrix
	matrix := [3][3]int{
		{1, 2, 4},
		{2, 6, 7},
		{3, 4, 5},
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}

	// create a 2d slice
	rows, cols := 3, 4
	matrix1 := make([][]int, 3)
	for i := 0; i < rows; i++ {
		matrix1[i] = make([]int, cols)
	}
	printmatrix(&matrix1)

	// Example of sorting an array
	arrToSort := []int{5, 3, 4, 1, 2}
	fmt.Println("Before sorting:", arrToSort)
	sort.Ints(arrToSort)
	fmt.Println("After sorting:", arrToSort)
}
