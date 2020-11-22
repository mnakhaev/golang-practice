package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second * 1)
			c1 <- "one"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 3)
			c2 <- "two"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			c3 <- "three"
		}
	}()

	// program has 15 seconds to execute (15 * 1 second from first function above)
	// every second it will receive value from c1 channel
	// every third second it will receive value from c2 channel
	// the order isn't strict - any of these two functions may be firstly executed
	// example: one, one, one, two, one, one, two, one, one

	// UPD: also added c3 wit hbigger delay
	for i := 0; i < 15; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received from channel #1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from channel #2:", msg2)
		case msg3 := <-c3:
			fmt.Println("Received from channel #3:", msg3)
		}
	}
}
