package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Println(e)
	fmt.Println(e.(customErr).info)

}

func main() {
	ce1 := customErr{
		info: "So sleepy!",
	}

	foo(ce1)
}
