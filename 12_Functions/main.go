package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	number := add(5, 6)
	fmt.Printf("The sum is: %d\n", number)
}
