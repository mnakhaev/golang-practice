package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
The main idea behind Drop Pattern is to have a limit on the amount of work that can be done at any given moment.

We have:
- a buffered channel that provides signaling semantic
- a number of worker goroutines
- a manager goroutine that:
	takes the work and sends it to the worker goroutine
	if there is more work than worker goroutines can process and buffered channel is full, manager goroutine will drop the work
*/

func main() {
	const capacity = 10
	ch := make(chan string, capacity)

	go func() {
		for i := range ch {
			fmt.Printf("Received element #%s\n", i)
		}
	}()

	const work = 200
	for i := 0; i < work; i++ {
		select {
		case ch <- strconv.Itoa(i):
			fmt.Printf("Sent task #%d to the channel...\n", i)
		default:
			fmt.Printf("Dropped task #%d\n", i)
		}
	}

	time.Sleep(time.Second)
}
