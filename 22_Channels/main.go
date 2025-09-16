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

func task(done chan string) {

	done <- "This is for the practice of channels"

}

func doneTask(done chan string) {
	msg := <-done
	fmt.Println("Now the channel send the message and done task recieved it", msg)
}

func main() {

	done := make(chan string)
	go task(done)
	go doneTask(done) // we need to pass channels as the parameters
	// waiting for the task to be done
	fmt.Println("Main function is done....")
	fmt.Scanln()
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
