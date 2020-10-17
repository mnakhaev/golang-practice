package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("Geometry:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main() {
	r := rect{
		width:  10,
		height: 20,
	}
	c := circle{radius: 100}
	measure(r)
	measure(c)

}
