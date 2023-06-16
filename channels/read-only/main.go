package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go read(ch)

	ch <- 1
	ch <- 2
	ch <- 3

	time.Sleep(time.Second)
}

// readonly channel way
func read(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
