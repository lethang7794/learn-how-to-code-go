package main

import "fmt"

func main() {
	f := func(s string) {
		fmt.Println("This is printed from an anonymous func, but assigned to a variable.")
		fmt.Println(s)

		n := 0
		n++
	}
	f("- But, why do we need this? 🙄")
	f("- Because we can do some interesting things with it!")
}
