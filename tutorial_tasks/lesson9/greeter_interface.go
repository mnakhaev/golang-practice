package main

import "fmt"

type greeter interface {
	greet(string) string
}

type russian struct{}
type american struct{}

// the functions for both types should have the same names as in interface
func (r *russian) greet(name string) string {
	return fmt.Sprint("Привет, ", name)
}

func (a *american) greet(name string) string {
	return fmt.Sprint("Hello, ", name)
}

// pass exact name of interface to the func
func sayHello(g greeter, name string) {
	fmt.Print(g.greet(name))
}

func main() {
	sayHello(&russian{}, "Павел")
	sayHello(&american{}, "John")
}
