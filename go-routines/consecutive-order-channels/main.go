package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go createTask("task 1", done)
	<-done
	go createTask("task 2", done)
	<-done
	go createTask("task 3", done)
	<-done

}

func createTask(text string, done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(text)
	done <- struct{}{}
}
