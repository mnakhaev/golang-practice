package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Printf("Starting main function...")
	var wg sync.WaitGroup

	t1 := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go makeSomethingConcurrently(i, &wg)
	}
	wg.Wait()
	totalTimeConcurrent := time.Since(t1).Seconds()

	t2 := time.Now()
	for i := 0; i < 10; i++ {
		makeSomethingConsequently(i)
	}
	totalTimeConsequent := time.Since(t2).Seconds()

	t3 := time.Now()

	wg.Add(1)
	// waitGroup should be passed to the goroutine in such case
	go func(wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			makeSomethingConcurrentlyInCycle(i)
		}
		wg.Done()
	}(&wg)
	wg.Wait()

	totalTimeChanda := time.Since(t3).Seconds()

	fmt.Printf(
		"Consequent code took %.2f seconds, concurrent - %.2f, chanda -  %.2f\n",
		totalTimeConsequent, totalTimeConcurrent, totalTimeChanda,
	)

	fmt.Printf("main function has finished")
}

func makeSomethingConcurrently(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	latency := rand.Intn(500) + 500 // emulate some action
	time.Sleep(time.Duration(latency) * time.Millisecond)
	fmt.Printf("- makeSomethingConcurrently: %d\n", num)
}

func makeSomethingConcurrentlyInCycle(num int) {
	latency := rand.Intn(500) + 500 // emulate some action
	time.Sleep(time.Duration(latency) * time.Millisecond)
	fmt.Printf("- makeSomethingConcurrentlyInCycle: %d\n", num)
}

func makeSomethingConsequently(num int) {
	latency := rand.Intn(500) + 500 // emulate some action
	time.Sleep(time.Duration(latency) * time.Millisecond)
	fmt.Printf("- makeSomethingConsequently: %d\n", num)
}
