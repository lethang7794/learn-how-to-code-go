package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	// x, y, z are declared without a corresponding initialization.
	// The compiler will assign zero-value to them.
}
