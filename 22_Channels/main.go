package main

import "fmt"

// func processNumber(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Received number: ", num)
// 		time.Sleep(time.Second * 1) // simulate processing time

// 	}
// }

//recieving data from channel

// func nums(result chan int, a int, b int) {
// 	numResult := a + b
// 	result <- numResult

// }

func task(done chan bool) {

	defer func() {
		done <- true
	}()
	fmt.Println("Task is done....")
}

func main() {

	done := make(chan bool)
	go task(done)
	<-done // waiting for the task to be done
	fmt.Println("Main function is done....")
	// 	result := make(chan int)
	// 	go nums(result, 10, 20)
	// 	resNum := <-result
	// 	fmt.Println(resNum)
	// channels are a way to communicate between go routines
	// numChan := make(chan int)
	// go processNumber(numChan)
	// for {
	// 	numChan <- rand.Intn(100) // blocking
	// }
	// buffered channel
	// messageChan := make(chan string)

	// messageChan <- "ping" // blocking
	// msg := <-messageChan
	// fmt.Println(msg)
}
