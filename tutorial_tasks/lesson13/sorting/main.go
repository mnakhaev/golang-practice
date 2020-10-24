package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	kids := []Person{
		{"Max", 9},
		{"Alan", 91},
		{"Tom", 63},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

}
