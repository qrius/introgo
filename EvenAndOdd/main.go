package main

import (
	"fmt"
)

func main() {
	numList := make([]int, 11)

	fmt.Println("length of slice NumList: ", len(numList))

	for i := range numList {
		numList[i] = i
		if numList[i]%2 != 0 {
			fmt.Println(numList[i], " is odd")
		} else {
			fmt.Println(numList[i], " is even")
		}
	}
	fmt.Println("numList: ", numList)

}
