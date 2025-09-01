package main

import "fmt"

//constants can be declared outside functions but variables cannot
const first int = 13

func main() {
	//constants

	const num int = 43

	fmt.Println("The value of num is:", num)
	fmt.Println("The value of first is:", first)

	// constants cannot be changed
	// num = 34 // this will throw an error
}
