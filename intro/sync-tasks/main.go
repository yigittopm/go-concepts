package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	createTask(100, "task1")
	createTask(200, "task2")
	createTask(0, "task3")
	createTask(100, "task4")

	fmt.Println("elapsed: ", time.Since(now))
}

func createTask(t time.Duration, text string) {
	time.Sleep(t * time.Millisecond)
	fmt.Println(text)
}
