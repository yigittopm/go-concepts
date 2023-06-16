package main

import "fmt"

func main() {
	ch := make(chan int)

	go write(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func write(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}