package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	var mu sync.Mutex

	go func() {
		mu.Lock()
		count = 10
		mu.Unlock()
	}()

	mu.Lock()
	count = 15
	mu.Unlock()

	time.Sleep(time.Second)

	mu.Lock()
	fmt.Println("count: ", count)
	mu.Unlock()
}
