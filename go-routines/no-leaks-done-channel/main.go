package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("go routines before work: ", runtime.NumGoroutine())

	done := make(chan struct{})

	go work(done, nil)

	go func() {
		time.Sleep(time.Second)
		fmt.Println("cancelling all go routines, listening on done")
		close(done)
	}()

	fmt.Println("number of goroutines while working: ", runtime.NumGoroutine())

	time.Sleep(2 * time.Second)

	fmt.Println("number of goroutines after closing done: ", runtime.NumGoroutine())
}

func work(done chan struct{}, strings <-chan string) {
	defer fmt.Println("work is done")

	for {
		select {
		case s := <-strings:
			fmt.Println(s)
		case <-done:
			return
		}
	}
}
