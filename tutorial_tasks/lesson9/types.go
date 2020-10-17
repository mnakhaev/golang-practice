package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	// Example of built-in type (anonymous field).
	// If define `Person Person`, then we won't be able to call c.Talk() directly
	Person
	Model string
}

type AndroidNew struct {
	Person Person
	Model  string
}

func main() {
	a := new(Person)
	a.Name = "Peter"
	a.Talk()

	b := new(Android)
	b.Name = "John"
	b.Person.Talk() // Smth like default way

	c := new(Android)
	c.Name = "Mark"
	c.Talk() // Possible only because of using built-in type

	d := new(AndroidNew)
	d.Person.Name = "Richard"
	// d.Name = "Denis" - impossible because d instance has no attribute 'Name'.
	// It's permitted to do only via `Person` structure
	d.Person.Talk()

}
