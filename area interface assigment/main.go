package main

import "fmt"

type shape interface {
	setName() string
	getArea() float64
}

type square struct {
	name       string
	sideLength float64
}

type triangle struct {
	name   string
	height float64
	base   float64
}

func main() {
	s := square{}
	s.sideLength = 5
	s.name = "square"

	t := triangle{}
	t.base = 2
	t.height = 3
	t.name = "triangle"

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println("The area of ", s.setName(), "is: ", s.getArea())
}

func (sq square) getArea() float64 {
	a := sq.sideLength * sq.sideLength
	return a
}

func (t triangle) getArea() float64 {
	a := 0.5 * t.base * t.height
	return a
}

func (sq square) setName() string {
	return sq.name
}
func (t triangle) setName() string {
	return t.name
}
