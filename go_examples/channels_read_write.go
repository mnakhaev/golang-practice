package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// write down values to channel in one goroutine
	go func() {
		for i := 1; i < 10; i++ {
			c <- i
			fmt.Println("Written value to the channel: ", i)
		}
		// close channel after cycle is finished to avoid deadlock
		close(c)
	}()

	// Read values from the channel
	for i := 1; i < 10; i++ { // or just `for {`
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println("1. Got value from the channel: ", msg)
	}

}
