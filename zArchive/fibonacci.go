// Fibonacci sequence for (n), recursively
package main

import (
	"fmt"
)

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	return (fib(n-1) + fib(n-2))

}

func main() {
	fmt.Println(fib(9))
}
