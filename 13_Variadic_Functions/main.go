package main

import "fmt"

func vardiacFunc(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {

	//passing slice as parameter in variadic function
	sliceNum := []int{5, 43, 32, 2}
	result := vardiacFunc(sliceNum...) //... is used to pass slice as parameter in variadic function
	fmt.Println(result)
}
