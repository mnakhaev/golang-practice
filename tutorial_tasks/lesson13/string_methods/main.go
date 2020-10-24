package main

import (
	"fmt"
	"strings"
)

func stringMethods() {
	fmt.Println(
		strings.Contains("test_string", "test"), // true

		strings.Count("test_string", "t"), //3

		strings.HasPrefix("test_string", "test"), // true

		strings.HasSuffix("test_string", "string"), // true

		strings.Index("test_string", "t"), // 0

		strings.Join([]string{"a", "b", "c"}, "-"), // a-b-c

		strings.Split("a-b-c", "-"), //  [a b c]

		strings.Repeat("test", 3), // testtesttest

		strings.Replace("test_string", "test", "new", 1), // new_string

		strings.ToLower("TEST"), // test

		strings.ToUpper("test"), // TEST

	)
}

func main() {
	stringMethods()
}
