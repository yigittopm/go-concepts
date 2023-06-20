package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)
	go fmt.Println("I try to print")

	for {
	}

}