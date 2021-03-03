package main

import "fmt"

type myInt int

var x myInt

func main() {
	fmt.Println("x:", x)
	fmt.Printf("Type of x: %T\n", x)

	x = 42
	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

}
