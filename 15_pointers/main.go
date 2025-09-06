package main

import "fmt"

// func nums(a int) {
// 	a = 10
// 	fmt.Println("Value of a in function:", a)
// }

func numVar(a *int) {
	*a = 10
	fmt.Println("Value of a in function:", *a)
}

func main() {

	//reference variable
	a := 5
	// nums(a)
	numVar(&a)
	fmt.Println("Value of a in main:", a)
}
