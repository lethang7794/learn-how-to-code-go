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

	var mu sync.Mutex

	for i := 0; i < numGoroutines; i++ {
		go func() {
			mu.Lock()
			c := counter
			// runtime.Gosched()
			fmt.Printf("%v\t", counter)
			c++
			counter = c
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println(runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println(runtime.NumGoroutine(), counter)
}
