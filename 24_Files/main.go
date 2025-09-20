package main

import (
	"fmt"
	"os"
)

func main() {
	f, error := os.Open("textFile.txt")
	if error != nil {
		panic(error)
	}

	fileinfo, error := f.Stat()
	if error != nil {
		panic(error)
	}
	fmt.Println("Your file name is: ", fileinfo.Name())
}
