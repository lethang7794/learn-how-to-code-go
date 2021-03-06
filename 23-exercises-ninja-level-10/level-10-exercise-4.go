package main

import (
	"fmt"
)

func main() {
	quit := make(chan bool)
	c := gen(quit)

	receive(c, quit)

	fmt.Println("about to exit")
}

func receive(c <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-c:
			fmt.Printf("Select from channel c: %v\n", v)
		case v := <-quit:
			fmt.Printf("Select from channel quit: %v\n", v)
			return
		}
	}
}

func gen(quit chan<- bool) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
		quit <- true
	}()

	return c
}
