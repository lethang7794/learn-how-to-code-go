package main

import "fmt"

func main() {
	returnFunc := aFunc()

	returnValueFromReturnFunc := returnFunc()

	fmt.Println(returnValueFromReturnFunc)
}

func aFunc() func() string {
	fmt.Println("This is aFunc.")

	return func() string {
		return "This string is returned from the return function."
	}
}
