package main

import (
	"fmt"
)

// Test exercise #1 from https://golangs.org/goroutines

// sourceGopher passes elems from array to channel
func sourceGopher(downstream chan string) {
	for _, elem := range [5]string{"a", "b", "b", "cc", "dd"} {
		downstream <- elem
	}
	close(downstream)
}

// handleGopher checks if there are duplicates. In such case - skip passing to next downstream
func handleGopher(upstream, downstream chan string) {
	previous := ""
	for item := range upstream {
		if item != previous {
			downstream <- item
			previous = item
		} else {
			fmt.Println("Already exists:", item)
		}

	}
	close(downstream)
}

// printGopher prints the result
func printGopher(upstream chan string) {
	for i := range upstream {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sourceGopher(ch1)
	go handleGopher(ch1, ch2)
	printGopher(ch2)
}
