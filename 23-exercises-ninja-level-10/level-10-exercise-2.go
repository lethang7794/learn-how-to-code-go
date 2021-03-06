package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c, 42)

	receive(c)
}

func send(c chan<- int, i int) {
	fmt.Printf("From send function:\n")
	fmt.Printf("\t\tc   (value): %v\n", c)
	fmt.Printf("\t\tc    (type): %T\n", c)
	c <- i
}

func receive(c <-chan int) {
	fmt.Printf("From receive function:\n\t\t<-c (value): %v\n", <-c)
	fmt.Printf("\t\tc   (value): %v\n", c)
	fmt.Printf("\t\tc    (type): %T", c)
}
