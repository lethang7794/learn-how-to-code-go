package main

import "fmt"

func main() {
	foo()
}

func foo() (int, error) {
	defer fmt.Println("This is deferred 1 - it's executed last")
	defer fmt.Println("This is deferred 2")
	defer fmt.Println("This is deferred 3")
	defer fmt.Println("This is deferred 4")
	defer fmt.Println("This is deferred 5")
	defer fmt.Println("This is deferred 6 - it's executed first")

	return fmt.Println("This is from return")
}
