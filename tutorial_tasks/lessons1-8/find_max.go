// Напишите функцию с переменным числом параметров, которая находит наибольшее число в списке.

package main

import "fmt"

func findMax(args ...int) int {
	if len(args) == 0 {
		return -1
	}

	max := 0
	for _, num := range args {
		if num > max {
			max = num
			fmt.Println("New max:", max)
		}
	}
	fmt.Println("max:", max)
	return max

}

func main() {
	findMax(23, 1, 5, 3, 5, 3, 3, 7)
}
