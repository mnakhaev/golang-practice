package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Result 1"
	}()

	// Timeout is less then sleep interval (2 seconds) => program will always finish by timeout
	select {
	case res1 := <-c1:
		fmt.Println("#1: Channel value:", res1)
	case <-time.After(time.Second * 1):
		fmt.Println("#1: Program finished by 1-second timeout")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Result 2"
	}()

	// This case can return timeout or value - depending on what will be executed firstly
	select {
	case res2 := <-c2:
		fmt.Println("#2: Channel value:", res2)
	case <-time.After(time.Second * 2):
		fmt.Println("#2: Program finished by 2-second timeout")
	}

	c3 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c3 <- "Result 3"
	}()

	// This case will always return channel value, because timeout is never reached
	select {
	case res3 := <-c3:
		fmt.Println("#3: Channel value:", res3)
	case <-time.After(time.Second * 3):
		fmt.Println("#3: Program finished by 3-second timeout")
	}
}
