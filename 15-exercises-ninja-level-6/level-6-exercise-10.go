package main

import "fmt"

func main() {
	incre := incrementor()
	fmt.Println(incre())
	fmt.Println(incre())
	fmt.Println(incre())
	fmt.Println(incre())
	fmt.Println(incre())

	incre2 := incrementor()
	fmt.Println(incre2())
	fmt.Println(incre2())
	fmt.Println(incre2())
}

func incrementor() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
