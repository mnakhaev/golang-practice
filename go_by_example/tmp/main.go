package main

import "fmt"

func main() {
	c1 := make(chan string, 1)
	c1 <- "msg"
	fmt.Println(<-c1)
}
