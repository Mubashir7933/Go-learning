package main

import "fmt"

func main() {
	//Arrays have same datatype and fixed size
	// var arr [5]int
	// fmt.Println(arr)

	// by default all values are initialized to zero
	// arr[0] = 10
	// fmt.Println(arr)

	//2d arrays
	var newArr [3][3]int = [3][3]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	fmt.Println(newArr)
}
