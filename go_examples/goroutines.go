package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// parseSmth emulates parsing of something
func parseSmth(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		latency := rand.Intn(500) + 500 // emulate some action
		time.Sleep(time.Duration(latency) * time.Millisecond)
		fmt.Println(name, ":", i)
	}
}

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	var wg sync.WaitGroup

	wg.Add(2)
	go parseSmth("goroutine №1", &wg)
	parseSmth("goroutine #2", &wg)

	wg.Wait()

	fmt.Printf("Parsing completed. Time elapsed: %.2f seconds\n", time.Since(t).Seconds())
	fmt.Println("done")
}
