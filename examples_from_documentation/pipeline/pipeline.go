package main

import (
	"fmt"
	"sync"
)

// based on https://blog.golang.org/pipelines

// gen writes some numbers to the channel
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out

	//go func() {
	//	for _, n := range nums {
	//		fmt.Println("Writing value in gen func:", n)
	//		out <- n
	//	}
	//	close(out)
	//}()
	//return out
}

// sq reads the numbers from inbound channel and writes their square values to the outbound
func sq(done <-chan struct{}, input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range input {
			select {
			case output <- n * n:
				fmt.Println("Writing square value:", n*n)
			case <-done:
				fmt.Println("Received DONE flag")
				return
			}
		}
	}()
	return output
}

// merge function converts a list of channels to a single channel by starting a goroutine for each input channel
// that copies the values to the single output channel.
// Then merge starts one more goroutine to close output channel after all sends are done
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	// ensure that all sends are done before closing the channel by using sync library
	var wg sync.WaitGroup
	out := make(chan int, 1)

	// Start an output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
				fmt.Println("Writing value to out channel:", n)
			case <-done:
				fmt.Println("Received done message!")
				return
			}
		}
		wg.Done() // mark channel as closed
	}
	wg.Add(len(cs)) // mark all channels as active

	for _, c := range cs {
		go output(c) // init the goroutine for ouput
	}

	// Start a goroutine to close `out` once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait() // makes something
		close(out)
	}()

	return out
}

func main() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits (e.g. in case of error), as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{}, 2)
	//defer close(done)

	input := gen(1, 2, 3, 4, 5)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(done, input)
	c2 := sq(done, input)

	// Consume the first value from the output
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9
	fmt.Println(<-out) // 4 or 9
	fmt.Println(<-out) // 4 or 9

	// No need in this 2 lines because of defer
	//done <- struct{}{}
	//done <- struct {}{}

	//for n := range merge(c1, c2) {
	//	fmt.Println(n)
	//}
}
