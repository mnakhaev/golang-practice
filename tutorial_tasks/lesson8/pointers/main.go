package main

import "fmt"

// This is an example of using interfaces with pointers

// Greeter interface contains greet method
type Greeter interface {
	greet()
}

type russian struct {
	name *string
}

type american struct {
	name *string
}

type newRussian struct {
	name string
}

// greet for russian type will change the name because the type uses pointer
func (r *russian) greet() {
	*r.name = "Russian name"
	fmt.Println("Russian name in function scope:", *r.name)
	fmt.Println("Privet,", *r.name)
}

// greet for american type will change the name because the type also uses pointer
func (a *american) greet() {
	*a.name = "American name"
	fmt.Println("American name in function scope:", *a.name)
	fmt.Println("Hello,", *a.name)
}

// greet for newRussian type won't change the name because no pointers used
func (n *newRussian) greet() {
	n.name = "New Russian name"
	fmt.Println("New Russian name in function scope:", n.name)
	fmt.Println("Hello,", n.name)
}

// sayHello accepts Greeter interface and calls `greet` method
func sayHello(g Greeter) {
	g.greet()
}

func main() {
	russianName := "Dima"
	sayHello(&russian{name: &russianName})                           // Russian name
	fmt.Println("Russian name after method is passed:", russianName) // Russian name

	americanName := "John"
	sayHello(&american{name: &americanName})                           // American name
	fmt.Println("American name after method is passed:", americanName) // American name

	newRussianName := "Pavel"
	sayHello(&newRussian{name: newRussianName})                             // New Russian name
	fmt.Println("New russian name after method is passed:", newRussianName) // Pavel (!)
}
