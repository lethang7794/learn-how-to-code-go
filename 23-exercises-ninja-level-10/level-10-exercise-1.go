package main

import (
	"fmt"
)

func main() {
	c := make(chan int) //* "A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type."
	// c <- 10 //! This won't work. WHY?
	go func() {
		c <- 10
	}()
	fmt.Println("c   :", c)
	fmt.Println("<-c :", <-c) //* Print the value received from the channel c

	c2 := make(chan int)
	go func() {
		c2 <- 20
	}()
	valueReceiveFromChannel := <-c2 //* Assign the value received from the channel c2 to a variable
	fmt.Println("-----")
	fmt.Println("<-c2:", valueReceiveFromChannel)

	c3 := make(chan int)
	go func() {
		c3 <- 30
	}()
	fmt.Println("-----")
	fmt.Println("<-c3:", <-c3)
	// fmt.Println(<-c3) //! This won't work. WHY?
	close(c3) //* "records that no more values will be sent on the channel"
	fmt.Println("<-c3:", <-c3, " (After c3 is closed)")

	c4 := make(chan int, 2)
	go func() {
		c4 <- 40
		c4 <- 41
		c4 <- 42
	}()
	fmt.Println("-----")
	fmt.Println(<-c4)
	fmt.Println(<-c4)
}
