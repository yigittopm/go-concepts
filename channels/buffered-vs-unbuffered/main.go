package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("buffered")
	buffered(2)

	fmt.Println("un-buffered")
	buffered(0)

	time.Sleep(time.Second)
	fmt.Println("done.")
}

func buffered(n int) {
	c := make(chan int, n)

	go func() {
		c <- 1
		fmt.Println("go routine 1")
	}()

	go func() {
		c <- 2
		fmt.Println("go routine 2")
	}()

	time.Sleep(time.Second)
	fmt.Println(<-c)

	time.Sleep(time.Second)
	fmt.Println(<-c)

}
