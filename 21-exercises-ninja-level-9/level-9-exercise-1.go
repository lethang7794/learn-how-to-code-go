package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())
	fmt.Println("Goroutines - Begin:\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("This is from the first goroutine.")
		wg.Done()
	}()

	fmt.Println("Goroutines:\t\t", runtime.NumGoroutine())

	go func() {
		fmt.Println("This is from the second goroutine.")
		wg.Done()
	}()

	fmt.Println("Goroutines:\t\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Goroutines - End:\t", runtime.NumGoroutine())
}
