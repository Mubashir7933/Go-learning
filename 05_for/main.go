package main

import "fmt"

// for -> it is the only keyword used for loops in go
func main() {
	// for using for as a while loop
	o := 10
	for o < 40 {

		fmt.Println("Value of o is:", o)
		o = o + 10
	}

	// classic for loop
	for i := 0; i <= 10; i++ {
		fmt.Println("Value of i is:", i)
	}

	//range in go new feature
	for k := range 11 {
		fmt.Println("Value of k is:", k)
	}

}
