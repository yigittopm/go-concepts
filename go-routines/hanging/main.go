package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// main go routine 1
	go func() {}() //  2
	go func() {}() //  3
	go func() {}() //  4

	fmt.Println("number of go routines", runtime.NumGoroutine())
	time.Sleep(time.Second)
	fmt.Println("main is done")
}
