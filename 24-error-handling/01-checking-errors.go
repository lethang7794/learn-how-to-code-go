package main

import "fmt"

func main() {
	n, err := fmt.Println("Trying to understand error handling in Go.")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
