package main

import "fmt"

func main() {
	var c chan int

	<-c

	for _ = range c {
		fmt.Println("I won't print")
	}

	c <- 1

	close(c)
}
