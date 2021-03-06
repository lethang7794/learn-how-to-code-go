package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i + 100
		}
		close(c2) // We need to close channel c2 so range know when to stop receive from channel c2
	}()

	for v := range c2 {
		fmt.Println(v)
	}
}
