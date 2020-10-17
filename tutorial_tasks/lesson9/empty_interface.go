package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]interface{}{} // interface{} is empty interface
	m["one"] = 1
	m["two"] = 2.0
	m["three"] = true

	for k, v := range m {
		switch v.(type) {
		case int:
			fmt.Println(k, "is integer")
		case float64:
			fmt.Println(k, "is float")
		default:
			fmt.Println(k, "is", reflect.TypeOf(v))
		}
	}
}
