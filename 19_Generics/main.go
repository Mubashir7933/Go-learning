package main

import "fmt"

// Generics used to pass any data type to a function or struct range over any data type

//

// using struct generics

type stack[T any] struct {
	elements []T
}

func main() {

	// using function generics
	myQueue := stack[int]{
		elements: []int{1, 2, 3},
	}

	//slicesNum := []int{1, 2, 3, 4, 5}
	fmt.Println(myQueue)

}
