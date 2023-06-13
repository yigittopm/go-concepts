package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	go write(ch)

	time.Sleep(1 * time.Second)

	for v := range ch {
		//time.Sleep(1 * time.Second)
		fmt.Println("read: ", v)
	}
}

func write(ch chan int) {
	for i := 1; i < 4; i++ {
		fmt.Println("writing: ", i)
		ch <- i
	}
	close(ch)
}
