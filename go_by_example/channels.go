package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//write down values to channel in one goroutine
	go func() {
		for i := 1; i < 10; i++ {
			c <- i
			fmt.Println("Written value to the channel: ", i)
		}
		// close channel after cycle is finished to avoid deadlock
		close(c)
	}()

	for i := 1; i < 10; i++ {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println("Got value from the channel: ", msg)
	}

	//Read values from channel
	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println("Got value from the channel: ", msg)
	}

}
