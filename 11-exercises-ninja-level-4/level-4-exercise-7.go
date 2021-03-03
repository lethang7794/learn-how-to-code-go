package main

import "fmt"

func main() {
	slice := [][]string{}
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)

	innerSlice1 := []string{"James", "Bond", "Shaken, not stirred"}
	innerSlice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	slice = append(slice, innerSlice1, innerSlice2)
	fmt.Println(slice)

	for i1, v1 := range slice {
		fmt.Println("Outer:", i1, v1)
		for i2, v2 := range v1 {
			fmt.Println("\tInner:", i2, v2)
		}
	}
}
