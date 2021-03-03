package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	numGoroutines := 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			c := counter
			runtime.Gosched()
			c++
			counter = c
			wg.Done()
		}()
		fmt.Println(runtime.NumGoroutine(), counter)
	}

	wg.Wait()
	fmt.Println(runtime.NumGoroutine(), counter)
}
