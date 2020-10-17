package main

import "fmt"

/*
1. Define main interface with sub-interfaces
2. For every sub-interface, define its method
3. Define needed structures
4. For every structure, define it's method with the same name as in step 2
5. Define new func with the same name as interface method. The func should accept the interface
6. Call needed method inside of this func

*/
type animalNew interface {
	walker
	runner
}

type bird interface {
	walker
	flier
}

type walker interface {
	walk()
}

type runner interface {
	run()
}

type flier interface {
	fly()
}

type cow struct{}
type eagle struct{}

func (c *cow) walk() {
	fmt.Println("Cow is walking...")
}

func (c *cow) run() {
	fmt.Println("Cow is running...")
}

func (e *eagle) walk() {
	fmt.Println("Eagle is walking...")
}

func (e *eagle) fly() {
	fmt.Println("Eagle is flying...")
}

func walkerFunc(w walker) {
	// this func is applicable for all of the structures
	w.walk()
}

func main() {
	var c animalNew = &cow{}
	walkerFunc(c)

	var e bird = &eagle{}
	walkerFunc(e)

}
