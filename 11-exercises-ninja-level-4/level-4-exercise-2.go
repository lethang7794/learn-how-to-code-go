package main

import "fmt"

func main() {
	slice := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 1010}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)
}
