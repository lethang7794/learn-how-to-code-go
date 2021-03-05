package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go foo(even, odd, quit)

	for {
		select {
		case v := <-even:
			fmt.Println("From even channel:", v)
		case v := <-odd:
			fmt.Println("From  odd channel:", v)
		case v := <-quit:
			fmt.Println("From quit channel:", v)
			return
		}
	}
}

func foo(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
