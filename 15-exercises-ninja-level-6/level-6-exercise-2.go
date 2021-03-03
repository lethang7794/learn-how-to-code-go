package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := foo(nums...)
	fmt.Println(sum)

	sum2 := bar(nums)
	fmt.Println(sum2)
}

func foo(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func bar(args []int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}
