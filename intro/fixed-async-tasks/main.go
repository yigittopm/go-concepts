package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	now := time.Now()

	go generateTask(done, "task1", 100)
	go generateTask(done, "task2", 200)
	go generateTask(done, "task3", 0)
	go generateTask(done, "task4", 100)

	<-done
	<-done
	<-done
	<-done

	fmt.Println("elapsed", time.Since(now))
}

func generateTask(done chan struct{}, text string, t time.Duration) {
	time.Sleep(t * time.Millisecond)
	fmt.Println(text)

	done <- struct{}{}
}
