package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(3)
	now := time.Now()
	go func() {
		createTask(1, "1")
		wg.Done()
	}()

	go func() {
		createTask(3, "2")
		wg.Done()
	}()

	go func() {
		createTask(2, "3")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("elapsed", time.Now().Sub(now))
	fmt.Printf("\n\n")

	now = time.Now()
	createTask(1, "1")
	createTask(3, "2")
	createTask(2, "3")
	fmt.Println("elapsed", time.Now().Sub(now))
}

func createTask(sec int, text string) {
	fmt.Println("before task: ", text)

	time.Sleep(time.Duration(sec) * time.Second)

	fmt.Println("after task: ", text)
}
