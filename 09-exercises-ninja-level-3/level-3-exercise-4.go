package main

import "fmt"

func main() {
	y := 1994

	for {
		if y > 2021 {
			break
		}
		fmt.Println(y)
		y++
	}
}
