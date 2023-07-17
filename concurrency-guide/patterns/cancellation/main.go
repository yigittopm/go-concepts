package main

import (
	"fmt"
	"time"
)

func main() {

	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()

		fmt.Println("doWork initiate...")

		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	fmt.Println("initiate...")
	d := <-terminated
	fmt.Println(d)
	fmt.Println("Done.")
}
