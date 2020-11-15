package main

import (
	"fmt"
)

func main() {

	messages := make(chan string, 1)
	signals := make(chan bool)

	// if value exists in messages, then it will be received. otherwise make as in default
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	// if value exists in messages, then it will be sent. otherwise make as in default
	msg := "Hello!"
	select {
	case messages <- msg: // works for bufferized channels! otherwise, will be done by default
		fmt.Println("Sent value", msg, "to messages channel")
	default:
		fmt.Println("No message sent")
	}

	// it's allowed to use multiple case statement
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("Nothing received")
	}
}
