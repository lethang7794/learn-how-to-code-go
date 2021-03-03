package main

import "fmt"

type myInt int

var x myInt
var y int

func main() {
	fmt.Println("Value of x:", x)
	fmt.Printf("Type of x: %T\n", x)

	x = 42
	fmt.Println("Value of x:", x)
	fmt.Printf("Type of x: %T\n", x)

	y = int(x)
	fmt.Println("Value of y:", y)
	fmt.Printf("Type of y: %T\n", y)
}
