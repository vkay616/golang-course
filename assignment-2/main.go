package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	sh1 := triangle{
		height: 10,
		base:   5,
	}

	sh2 := square{
		sideLength: 4,
	}

	printArea(sh1)
	printArea(sh2)
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}

func (sq square) getArea() float64 {
	return (sq.sideLength * sq.sideLength)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
