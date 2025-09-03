package main

import (
	"fmt"
	"slices"
)

// Slices are basically dynamic arrays
// When the size of the array is not known at the time of declaration
// Slices are built on top of arrays
func main() {
	//

	// when you reach maximum capacity, a new array is created with double the capacity

	//var nums = make([]int, 3, 5) // now here the capacity is 5
	// fmt.Println(len(nums))
	// nums = append(nums, 23)
	// nums = append(nums, 22)
	// nums = append(nums, 20)
	// nums = append(nums, 24)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	//copy function
	// copies from source to destination
	var test = make([]int, 3) // here the capacity is also 3
	test = append(test, 3)
	var test2 = make([]int, len(test))

	copy(test2, test)
	fmt.Println(test2)

	// slices operator
	// var number = []int{2, 3, 4, 5, 6, 7, 8}
	// fmt.Println(number[1:4]) // 4 is excluded
	// fmt.Println(number[:4])  // from start to 4 excluded
	// fmt.Println(number[1:])  // from 1 to end
	// fmt.Println(number[:])   // from start to end

	//slices equal package

	var s1 = []int{1, 2, 3}
	var s2 = []int{1, 2, 3}

	fmt.Println(slices.Equal(s1, s2))

	//2d slices
	var NewSlice = [][]int{{1, 2, 3}, {9, 8, 7}}
	fmt.Println(NewSlice)
}
