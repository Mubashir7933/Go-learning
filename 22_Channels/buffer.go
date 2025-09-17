package main

import "fmt"

func main() {
	// mainCh := make(chan int, 3) // buffered channel with capacity of 2
	// mainCh <- 1
	// mainCh <- 2
	// fmt.Println("This is testing for buffered channels")
	// fmt.Println("This is testing for buffered channels", <-mainCh)
	// fmt.Println("This is testing for buffered channels", <-mainCh)
	// mainCh <- 3
	// fmt.Println("Let's see if it's running")

	// inside the goroutine channels using buffered channels you can pass and recieve value without blocking in the main function however using unbuffered channels will block the main function until the value is recieved

	channelMain := make(chan int)
	go func() {

		channelMain <- 42
		fmt.Println("This is inside the goroutine")
	}()
	msg := <-channelMain

	fmt.Println("the value of goroutine", msg)
}
