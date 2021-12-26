package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
The main idea behind Fan Out Pattern is to have:

- a channel that provides a signaling semantics
- channel can be buffered, so we don't wait on immediate receive confirmation
- a goroutine that starts multiple (other) goroutines to do some work
- a multiple goroutines that do some work and use signaling channel to signal that the work is done
*/

func main() {
	employeesNum := 3
	ch := make(chan string, employeesNum)

	for i := 1; i <= employeesNum; i++ {
		i := i
		go func() {
			fmt.Printf("Starting some work in goroutine #%d...\n", i)
			latency := time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(latency)
			ch <- fmt.Sprintf("work is done in goroutine #%d with latency %v", i, latency)
		}()
	}

	for employeesNum > 0 {
		result := <-ch
		fmt.Println("received result:", result)
		employeesNum--
	}

}
