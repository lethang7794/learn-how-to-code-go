package main

import (
	"fmt"
	"sync"
)

func main() {
	// SOLUTION 1
	c := make(chan string)

	go func() {
		numGoroutines := 10
		for i := 0; i < numGoroutines; i++ { //! for statement reuse i for each iteraction
			i := i
			go func() {
				for j := 0; j < 10; j++ {
					c <- fmt.Sprintf("i:%v - j: %v", i, j)
				}
			}()
		}
	}()

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
	//! We can't range over the channel here. WHY?

	// SOLUTION 2: Range over the channel
	c2 := make(chan int)
	var wg sync.WaitGroup

	go func() {
		numGoroutines := 10
		for i := 0; i < numGoroutines; i++ {
			i := i //! Get "a fresh version of the variable with the same name" - https://golang.org/doc/effective_go#channels
			wg.Add(1)
			go func() {
				for j := 0; j < numGoroutines; j++ {
					c2 <- j*10 + i
				}
				wg.Done()
			}()
		}

		wg.Wait()
		close(c2)
	}()

	for v := range c2 {
		fmt.Println(v)
	}
}
