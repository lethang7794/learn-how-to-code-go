package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is printed from an anonymous func, which are invoked right after declared.")
		fmt.Println("----------")
	}()

	func(s string) {
		fmt.Println("This is printed from another anonymous func, which has a parameter.")
		fmt.Println("This is the argument received:", s)
		fmt.Println("----------")
	}("ABC XYZ")
}
