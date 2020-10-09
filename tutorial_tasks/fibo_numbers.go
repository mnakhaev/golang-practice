// Последовательность чисел Фибоначчи определяется как fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
// Напишите рекурсивную функцию, находящую fib(n).

package main

import (
	"fmt"
)

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-2) + fibo(n-1)
	}
}

func main() {
	fmt.Println(fibo(5))
}
