package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int

	bound := 1000
	wg.Add(bound)

	for i := 0; i < bound; i++ {
		go func() {
			defer wg.Done()

			var localCount int
			mu.Lock()
			localCount = count + 1
			count = localCount
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("count: ", count)
}
