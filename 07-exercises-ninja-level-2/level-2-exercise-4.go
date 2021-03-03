package main

import "fmt"

func main() {
	a := 32
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%x\n", a)

	b := a << 1
	fmt.Printf("%d\n", b)
	fmt.Printf("%b\n", b)
	fmt.Printf("%x\n", b)
}
