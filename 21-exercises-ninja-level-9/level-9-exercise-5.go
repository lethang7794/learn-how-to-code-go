package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	numGoroutines := 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println(runtime.NumGoroutine(), "\t", atomic.LoadInt64(&counter))
	}

	wg.Wait()
	fmt.Println(runtime.NumGoroutine(), "\t", atomic.LoadInt64(&counter))
}
