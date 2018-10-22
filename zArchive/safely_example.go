package main

import (
	"errors"
	"time"

	"github.com/qrius/introgo/safely"
)

func message() {
	println("Inside goroutine")
	panic(errors.New("Oops!"))

}

func main() {
	safely.Go(message)
	println("outside goroutine")
	time.Sleep(1000)
}
