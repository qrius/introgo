/*
Package Name: FizzBuzz
Purpose: Print numbers from 1 to 100, but for multiple of 3, print "Fizz" instead of the number, for multiple of 5, print "Buzz".  For multiple of both 3 and 5, print "FizzBuzz"
Example: 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, ... 13, 14, FizzBuzz, ...

Algorithm

for i from 1 to 100 do
	if remainder of ((i/3) and (i/5)) then print FizzBuzz
	elif remainder of (i/3) = 0 then print Fizz
	elif remainder of (i/5) = 0 then print Buzz
	else print i
*/
package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Printf("%s ", "FizzBuzz")
		} else if (i % 3) == 0 {
			fmt.Printf("%s ", "Fizz")
		} else if i%5 == 0 {
			fmt.Printf("%s ", "Buzz")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}
