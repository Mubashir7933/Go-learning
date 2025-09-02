package main

import "fmt"

//Slices are basically dynamic arrays
// When the size of the array is not known at the time of declaration
// Slices are built on top of arrays
func main() {
	var mySlice []int
	fmt.Println(mySlice == nil)

	var newSlice = make([]int, 4)
	fmt.Println(newSlice == nil)
}
