package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var count int

	wg.Add(1)
	go func() {
		defer wg.Done()
		count++
	}()
	wg.Wait()
	fmt.Println(count)
}
