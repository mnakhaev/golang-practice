package main

import "fmt"

// ping receives `pings` channel only for send the values
// directions can be removed, they are for visibility and types security
func ping(pings chan<- string, msg string) {
	pings <- msg // `pings` channel accepts message
}

// pong receives `pings` channel for accepting values and `pongs` channel for sending values
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings                   // `pongs` channel creates a new `msg` variable, which is received from `pings`
	msg = "modified message: " + msg // then message is modified
	pongs <- msg                     // then this message is passed to `pongs`
}

func main() {
	pings := make(chan string, 1) // using non-buffered channel causes fatal error
	pongs := make(chan string, 1) // pings channel accepts message
	ping(pings, "Passed message") // create new message and pass it to `pings` channel
	pong(pings, pongs)            // accept this message and pass it to `pongs`
	fmt.Println(<-pongs)
}
