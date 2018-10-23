package main

import "fmt"

func swap(xPtr *int, yPtr *int) {
	println("parameter y:", yPtr)
	println("address of yPtr: ", &yPtr)
	println("content of *yPtr: ", *yPtr)
	//set c to the content at the address dereferenced by *yPtr
	c := *yPtr
	*yPtr = *xPtr
	*xPtr = c
}

func main() {
	x := 1
	y := 2
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	println("address of y: ", &y)

	swap(&x, &y)

	fmt.Println("x is now = ", x)
	fmt.Println("y is now = ", y)
}
