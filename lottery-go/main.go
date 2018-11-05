package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// lottery := newDraw()

	var guess int

	fmt.Println("enter a number between 0-9")
	fmt.Scanf("%d", &guess)

	fmt.Println("Your guess : ", guess)

	// rand.Seed(time.Now().UnixNano())
	fmt.Println("Here's a random number between 0 -9 : ", randomInt(0, 10)) //get an int in the 1...10 range

}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
