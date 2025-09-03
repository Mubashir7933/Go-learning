package main

import "fmt"

func main() {

	//range with integer

	// nums := []int{1, 2, 3, 4, 5}

	// // looping through slice using range
	// for i, num := range nums {
	// 	fmt.Printf("Index is: %d,value is %d\n", i, num)
	// }

	//range with string
	// name := "mubashir"
	// for s, m := range name {
	// 	fmt.Printf("The index is %d and the value is %c\n", s, m)

	// %c is used to print character
	// %d is used to print integer

	//range with map
	newMap := map[string]int{
		"Syed":   23,
		"Ali":    25,
		"Haseeb": 26,
		"Veysel": 24,
	}
	for k, v := range newMap {
		fmt.Printf("The key is %s and the value is %d\n", k, v)
	}

	// %s is used to print string
	// %d is used to print integer

}
