// Напишите программу, которая меняет местами два числа (x := 1; y := 2; swap(&x, &y) должно дать x=2 и y=1)

package main

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
}
