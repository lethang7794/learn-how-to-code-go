package main

import "fmt"

func main() {
	a := 2
	b := 2

	switch {
	case a < b:
		fmt.Println("a < b")
	case a > b:
		fmt.Println("a > b")
	default:
		fmt.Println("a = b")
	}
}
