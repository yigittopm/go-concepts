package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// go run -race main.go
func main() {
	var count int32
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1e5; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1)
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("count: ", count)
}
