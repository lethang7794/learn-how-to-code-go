package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := sum(nums...)
	fmt.Println(s)

	sEven := sumEven(sum, nums...)
	fmt.Println(sEven)

	sOdd := sumOdd(sum, nums...)
	fmt.Println(sOdd)
}

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func sumEven(sum func(...int) int, nums ...int) int {
	evens := []int{}

	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}

	s := sum(evens...)
	return s
}

func sumOdd(sum func(...int) int, nums ...int) int {
	odds := []int{}

	for _, num := range nums {
		if num%2 == 1 {
			odds = append(odds, num)
		}
	}

	s := sum(odds...)
	return s
}
