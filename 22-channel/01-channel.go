package main

import "fmt"

func main() {
	c := make(chan int)

	// Not working
	// c <- 42

	go func() {
		c <- 42
		c <- 43
		c <- 44
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	c2 := make(chan int, 1)
	go func() {
		c2 <- 100
		c2 <- 101
		c2 <- 102
	}()

	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
