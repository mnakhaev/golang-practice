package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	for i := 1; i < 10; i++ {
		go writeData(i, c)
	}

	for i := 1; i < 10; i++ {
		// reading is happening only if there is something to read.
		// if channel is empty, then `main` goroutine waits until `writeData` writes something there
		fmt.Println("Got value from the channel:", <-c)
	}

}

func writeData(i int, c chan int) {
	// place with print makes sense. it may mislead is placed after reading from channel
	fmt.Println("Started writing value to the channel: ", i)
	c <- i
	fmt.Println("Finished writing value to the channel: ", i)
}
