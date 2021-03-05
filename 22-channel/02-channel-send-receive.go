package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	// The program will exit before run this go routine
	go bar(c)

	bar(c)

	fmt.Println("About to exit!")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
