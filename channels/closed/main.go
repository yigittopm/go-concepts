package main

import "fmt"

func main() {
	c := make(chan int)
	// this gets ignored as well
	close(c)
	<-c

	for _ = range c {
		fmt.Println("I won't print")
	}

	c <- 1
}
