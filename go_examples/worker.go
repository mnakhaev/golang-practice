package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

// process func receives data from channel for a worker
func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("Worker %d received %d\n", w.id, data)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	// c := make(chan int, 20)
	c := make(chan int)
	for i := 0; i < 5; i++ { // i is number of workers
		worker := Worker{id: i}
		go worker.process(c) // execution time depends on number in channel
	}

	for {
		select {
		case c <- rand.Int():
			fmt.Println("Generated new value!", <-c)
		case t := <-time.After(time.Millisecond * 100): // timeout is another method
			fmt.Println("Timeout after", t)
		default:
			fmt.Println("Skipping value...")
		}
		time.Sleep(time.Millisecond * 50)
	}
}
