package main

import "fmt"

func main() {
	y := 2021

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(&y)
	fmt.Printf("%T\n", &y)
}
