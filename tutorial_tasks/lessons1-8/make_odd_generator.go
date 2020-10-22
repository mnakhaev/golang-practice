// Напишите функцию makeOddGenerator, генерирующую нечётные числа

package main

import "fmt"

func oddGenerator() func() uint {
	i := uint(1)
	return func() (result uint) {
		result = i
		i += 2
		return result
	}
}

func main() {
	nextOdd := oddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

}
