package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	ch <- "value"
	val, ok := <-ch
	fmt.Println("val: ", val)
	fmt.Println("ok: ", ok)

	close(ch)

	val, ok = <-ch
	fmt.Println("val: ", val)
	fmt.Println("ok: ", ok)
}
