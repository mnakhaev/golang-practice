package main

import (
	"fmt"
	"strings"
)

// Test exercise #2 from https://golangs.org/goroutines

func mainGopher(downstream chan string, sentences []string) {
	for _, sentence := range sentences {
		fmt.Println("\nCurrent sentence:", sentence)
		res := strings.Fields(sentence)
		for _, r := range res {
			fmt.Println("\tSending current word to channel:", r)
			downstream <- r
		}
	}
	close(downstream)
}

func destGopher(upstream, downstream chan string) {
	fmt.Println("\nDestination Gopher is in progress...")
	for word := range upstream {
		downstream <- word
	}
	close(downstream)
}

func printingGopher(upstream chan string) {
	fmt.Println("\nPrinting Gopher is in progress...")
	for word := range upstream {
		fmt.Println("Received word:", word)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	sentences := []string{
		"London is the capital of Great Britain",
		"Hello, World",
		"This is test example",
	}
	go mainGopher(ch1, sentences)
	go destGopher(ch1, ch2)
	printingGopher(ch2)
	//time.Sleep(time.Second * 1)
}
