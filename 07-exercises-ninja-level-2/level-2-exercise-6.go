package main

import "fmt"

const (
	_  = iota
	y1 = iota + 2021
	y2 = iota + 2021
	y3 = iota + 2021
	y4 = iota + 2021
)

func main() {
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
