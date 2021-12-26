package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
The main idea behind Wait For Task pattern is to have:

- a channel that provides a signaling semantics
- a goroutine that waits for task, so it can do some work
- a goroutine that sends work to the previous goroutine
*/

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("213\n")
		result := <-ch
		fmt.Printf("Starting execution of the task: %s\n", result)
	}()

	// simulate some work
	latency := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(latency)

	task := "test task"
	ch <- task
	fmt.Printf("Sent task %q\n", task)
	wg.Wait()
}
