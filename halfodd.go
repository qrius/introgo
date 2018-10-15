/*
Package halfodd taks an integer and halves it and returns true if it was env or false if it was odd.  For example, half(1) should return (0, false) and half(2) shoud return (1, true)

*/

package main

import (
	"fmt"
	"log"
)

func evenOrNot(x int) string {

	if x/2%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}

}

func main() {
	var x int

	fmt.Println("enter a number: ")
	if _, err := fmt.Scan(&x); err != nil {
		log.Print("  Scan for x failed, due to ", err)
		return
	}

	h := evenOrNot(x)

	fmt.Printf("Half of %v is an %s number", x, h)

}
