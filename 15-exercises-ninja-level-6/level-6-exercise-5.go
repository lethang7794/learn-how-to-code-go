package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	a float64
}

type CIRCLE struct {
	radius float64
}

type SHAPE interface {
	area() float64
}

func (sq SQUARE) area() float64 {
	return sq.a * sq.a
}

func (c CIRCLE) area() float64 {
	pi := math.Pi
	return 2 * pi * c.radius
}

func info(s SHAPE) {
	fmt.Println(s.area())
}

func main() {
	square := SQUARE{
		a: 2,
	}

	circle := CIRCLE{
		radius: 3,
	}

	// squareArea := square.area()
	// circleArea := circle.area()
	// fmt.Println(squareArea)
	// fmt.Println(circleArea)

	info(square)
	info(circle)
}
