package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var count int

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			var localCount int
			localCount += count + i + 1
			fmt.Println("local count", localCount)
		}(i)
	}
	wg.Wait()
}
