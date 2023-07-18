package main

import (
	"fmt"
	"sync"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	wg.Add(2)
	register(cond, func() {
		defer wg.Done()
		fmt.Println("worker 1")
	})

	register(cond, func() {
		defer wg.Done()
		fmt.Println("worker 2")
	})

	cond.Broadcast()
	wg.Wait()
}

func register(cond *sync.Cond, fn func()) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		cond.L.Lock()
		cond.Wait()
		fn()
		cond.L.Unlock()
	}()
	wg.Wait()
}
