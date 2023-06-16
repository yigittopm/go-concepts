package main

import (
	"fmt"
	"unsafe"
)

func main() {
	emptyStruct := make(chan struct{})
	intChan := make(chan int)

	go func() {
		emptyStruct <- struct{}{}
		intChan <- 1
	}()

	fmt.Println(unsafe.Sizeof(<-emptyStruct))
	fmt.Println(unsafe.Sizeof(<-intChan))
}
