package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	go createTastk("task 1")
	go createTastk("task 2")
	go createTastk("task 3")

	time.Sleep(time.Second)
}

func createTastk(text string) {
	fmt.Println(text)
}
