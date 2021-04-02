package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file-is-here")
	if err != nil {
		// fmt.Println(err) // Println formats using the default formats for its operands and writes to standard output.
		// log.Println(err) // Println calls Output to print to the standard logger.
		// log.Fatal(err)   // Fatal is equivalent to Print() followed by a call to os.Exit(1).
		panic(err) // When a function F calls panic, normal execution of F stops immediately...This termination sequence is called panicking and can be controlled by the built-in function recover.
	}

	fmt.Println("About to exit")
}
