package main

import "fmt"

func main() {
	arr := [5]int{11, 22, 33, 44, 55}

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Printf("%T", arr)
}
