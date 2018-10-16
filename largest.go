/*
largest is a function with one varidic parameter that finds the greatest number in a list of numbers
*/
package main

import (
	"fmt"
	"sort"
)

func maxOf(args ...int) int {
	sort.Ints(args)
	//fmt.Println("Args:", args)
	return args[len(args)-1]
}

func main() {
	xs := []int{1, 2, 8, 4, 45, 6, 7}
	fmt.Println("largest number is", maxOf(xs...))

}
