package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := make(chan int, 1)

	go func() {
		fmt.Println("writing 1")
		numbers <- 1

		fmt.Println("writing 2")
		numbers <- 2

		fmt.Println("writing 3")
		numbers <- 3

		fmt.Println("writing 4")
		numbers <- 4
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("read 1: ", <-numbers)
	fmt.Println("read 2: ", <-numbers)
	fmt.Println("read 3: ", <-numbers)
	fmt.Println("read 4: ", <-numbers)
}
