package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	<-ch
	<-ch

	fmt.Println("reads on closed channel are ignored")

	ch <- 1
}
