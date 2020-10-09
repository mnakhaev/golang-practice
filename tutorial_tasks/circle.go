package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func circleArea(c *Circle) float64 {
	// If I need to change global variables inside of function scope, then need to use pointers
	// like c.x = 2, c.y = 3, c.r = 5
	// without pointers, the values will be returned after exiting function scope
	fmt.Println("Circle data:", c.x, c.y, c.r)
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.x2, r.x1, r.y2)
	w := distance(r.x1, r.x2, r.x1, r.y1)
	fmt.Println("Length:", l, "Width:", w)
	return l * w
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	c.x = 10 // values in the structure will be updated
	c.y = 4
	fmt.Println(circleArea(&c)) // Example #1
	fmt.Println(c.area())       // Example #2

	r := Rectangle{10, 2, 100, 9}
	fmt.Println(r.area())
}
