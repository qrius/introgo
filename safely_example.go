package main

import (
	"github.com/qrius/introgo/safely"
	"errors"
	"time"
)

func message () {
	println("Inside goroutine")
	panic(errors.New("Oops!"))

}

func main (){
	safely.go(message)
	println("outside goroutine")
	time.sleep(1000)
}