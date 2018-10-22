package main

import "fmt"

func swap(xPtr *int, yPtr *int) {
	c := yPtr
	yPtr = xPtr
	xPtr = c
}

func main() {
	x := 1
	y := 2
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

	swap(&x, &y)

	fmt.Println("x is now = ", x)
	fmt.Println("y is now = ", y)
}
