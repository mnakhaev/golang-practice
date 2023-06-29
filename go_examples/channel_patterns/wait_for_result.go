package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
The main idea behind Wait For Result pattern is to have:

- a channel that provides a signaling semantics
- a goroutine that does some work
- a goroutine that waits for that work to be done
*/

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	defer wg.Done()

	wg.Add(1)
	go func() {
		fmt.Println("Starting some work...")
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "The work is done"
		wg.Wait()
	}()

	result := <-ch
	fmt.Println("Received result:", result)
}
