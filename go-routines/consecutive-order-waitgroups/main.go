package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go createTask("task 1")
	wg.Wait()

	wg.Add(1)
	go createTask("task 2")
	wg.Wait()

	wg.Add(1)
	go createTask("task 3")
	wg.Wait()
}

func createTask(text string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(text)
	wg.Done()
}
